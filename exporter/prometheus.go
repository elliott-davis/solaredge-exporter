package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/elliott-davis/solaredge-go/solaredge"
	"time"
)

// Describe - loops through the API metrics and passes them to prometheus.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.APIMetrics {
		ch <- m
	}

}

// Collect function, called on by Prometheus Client library
// This function is called when a scrape is performed on the /metrics page
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	sevenPM := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, now.Location())
	sevenAM := sevenPM.Add(time.Hour * 12)
	if now.After(sevenPM) && now.Before(sevenAM) {
		siteOverview := solaredge.SiteOverview{}
		siteOverview.CurrentPower.Power = 0.0
		err := e.processMetrics(siteOverview, ch)
		if err != nil {
			log.Error("Error Processing Metrics", err)
			return
		}
		log.Info("Bailing out of metrics collection due to night time")
		return
	}
	// Scrape the Data from Solaredge
	client := solaredge.NewClient(nil, e.Config.APIToken)
	data, err := client.Site.Overview(e.Config.Site)
	if err != nil {
		log.Errorf("Error gathering data from SolarEdge %v", err)
	}

	if err != nil {
		log.Errorf("Error gathering Data from remote API: %v", err)
		return
	}

	// Set prometheus gauge metrics using the data gathered
	err = e.processMetrics(data, ch)

	if err != nil {
		log.Error("Error Processing Metrics", err)
		return
	}

	log.Info("All Metrics successfully collected")

}
