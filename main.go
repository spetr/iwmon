package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	configLoad("iwmon.yml")
	confRuntimeLoad()
	handleSigHup()

	r := prometheus.NewRegistry()
	//	r.MustRegister(prometheus.NewGoCollector())

	if conf.API.Prometheus {
		http.Handle("/metrics", promhttp.HandlerFor(
			r,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
	}

	go monIceWarpVersionUpdate(r)
	go monFsMailUpdate(r)

	log.Fatal(http.ListenAndServe(conf.API.Listen, nil))
}
