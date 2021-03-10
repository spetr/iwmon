package main

import (
	"fmt"
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
		var (
			iwResponse map[string]string
			err        error
		)
		for {
			if iwResponse, err = iwToolGet("system", "c_version", "c_settingsversion", "c_date"); err != nil {
				fmt.Printf("TODO - ERROR: %s", err.Error())
				time.Sleep(10)
			}

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
