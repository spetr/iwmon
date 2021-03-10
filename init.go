package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init")

	// Add Go module build info.
	//	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}
