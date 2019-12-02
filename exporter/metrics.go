package exporter

import (
	"github.com/elliott-davis/solaredge-go/solaredge"
	"github.com/prometheus/client_golang/prometheus"
)

func AddMetrics() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["Power"] = prometheus.NewDesc(
		prometheus.BuildFQName("solaredge", "site", "power"),
		"Current power output",
		[]string{}, prometheus.Labels{},
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data solaredge.SiteOverview, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(e.APIMetrics["Power"], prometheus.GaugeValue, data.CurrentPower.Power)

	return nil
}
