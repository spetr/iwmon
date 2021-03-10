package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	r := prometheus.NewRegistry()
	r.MustRegister(monIceWarpVersion)
	r.MustRegister(prometheus.NewGoCollector())

	http.Handle("/metrics", promhttp.HandlerFor(
		r,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
