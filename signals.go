package main

import (
	"os"
	"os/signal"
	"syscall"
)

func handleSigHup() {
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGHUP)
		for {
			<-sigs
			configLoad("iwmon.yml")
			confRuntimeLoad()
		}
	}()
}
