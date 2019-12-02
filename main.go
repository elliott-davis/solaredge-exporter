package main

import (
	"github.com/fatih/structs"
	conf "github.com/elliott-davis/solaredge-exporter/config"
	"github.com/elliott-davis/solaredge-exporter/exporter"
	"github.com/elliott-davis/solaredge-exporter/http"
	"github.com/infinityworks/go-common/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	log            *logrus.Logger
	applicationCfg conf.Config
	mets           map[string]*prometheus.Desc
)

func init() {
	applicationCfg = conf.Init()
	mets = exporter.AddMetrics()
	log = logger.Start(&applicationCfg)
}

func main() {
	log.WithFields(structs.Map(applicationCfg)).Info("Starting Exporter")

	exp := exporter.Exporter{
		APIMetrics: mets,
		Config:     applicationCfg,
	}

	http.NewServer(exp).Start()
}
