package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spetr/go-zabbix-sender"
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
	var (
		iwResponse map[string]string
		err        error
	)
	for {
		if iwResponse, err = iwToolGet("system", "c_version", "c_settingsversion", "c_date"); err != nil {
			fmt.Printf("TODO - ERROR: %s", err.Error())
			time.Sleep(10 * time.Second)
		}

		// Prometheus Exporter
		if conf.API.Prometheus {
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
		}

		// Zabbix Sender
		if conf.ZabbixSender.Enabled {
			var (
				metrics []*zabbix.Metric
				t       = time.Now().Unix()
			)
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "iw.os", runtime.GOOS, true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "iw.version", iwResponse["c_version"], true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "iw.settingsversion", iwResponse["c_settingsversion"], true, t))
			metrics = append(metrics, zabbix.NewMetric(conf.ZabbixSender.Hostname, "iw.date", iwResponse["c_date"], true, t))
			for i := range conf.ZabbixSender.Servers {
				z := zabbix.NewSender(conf.ZabbixSender.Servers[i])
				z.SendMetrics(metrics)
			}
		}

		time.Sleep(time.Duration(conf.IceWarp.Refresh.Version) * time.Second)
	}
}
