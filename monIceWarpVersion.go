package main

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	monIceWarpVersion = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "iw_version",
			Help: "IceWarp version",
		},
	)
)

// IceWarp version (from tool.sh)
func monIceWarpVersionUpdate(r *prometheus.Registry) {
	go func(r *prometheus.Registry) {
		for {
			iwResponse, _ := iwToolGet("system", "c_version", "c_settingsversion", "c_date")
			r.Unregister(monIceWarpVersion)
			monIceWarpVersion = prometheus.NewGauge(
				prometheus.GaugeOpts{
					Name: "iw_version",
					Help: "IceWarp version",
					ConstLabels: prometheus.Labels{
						"os":              runtime.GOOS,
						"version":         iwResponse["c_version"],
						"settingsversion": iwResponse["c_settingsversion"],
						"date":            iwResponse["c_date"],
					},
				},
			)
			r.MustRegister(monIceWarpVersion)
			monIceWarpVersion.Set(1)
			time.Sleep(conf.IceWarp.Refresh.Version * time.Second)
		}
	}(r)
}
