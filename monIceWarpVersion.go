package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	monIceWarpVersion = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        "iw_version",
			Help:        "IceWarp version",
			ConstLabels: prometheus.Labels{"version": "unknown", "os": "unknown"},
		},
	)
)

func monIceWarpVersionUpdate() {
	iwToolGet("system", "c_version")
	monIceWarpVersion = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        "iw_version",
			Help:        "IceWarp version",
			ConstLabels: prometheus.Labels{"version": "unknown", "os": "unknown"},
		},
	)
	monIceWarpVersion.Set(1)
}
