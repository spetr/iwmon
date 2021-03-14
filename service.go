package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spetr/service"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Started in terminal.")
	} else {
		logger.Info("Started with service manager.")
	}
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) run() {
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
	if conf.IceWarp.SNMP.Enabled {
		go monIceWarpSNMPUpdate(r)
	}
	go monFsMailUpdate(r)

	go func() {
		if err := http.ListenAndServe(conf.API.Listen, nil); err != nil {
			logger.Error(err.Error())
		}
	}()

	<-p.exit
}

func (p *program) Stop(s service.Service) error {
	logger.Info("Stopping")
	close(p.exit)
	return nil
}
