package exporter

import (
	"github.com/elliott-davis/solaredge-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	config.Config
}
