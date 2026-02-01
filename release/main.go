package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MA3CIN/WasmPulse/release/collector"
	"github.com/MA3CIN/WasmPulse/release/discovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	
	pidCollector := collector.NewPidCollector()
	reg.MustRegister(pidCollector)

	go discoverPIDsLoop(pidCollector)

	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)

	fmt.Println("Server starting on :8080/metrics")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// Continuous loop to find new PIDs.
func discoverPIDsLoop(c *collector.PidCollector) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		procInfoArr := discovery.DiscoverWASM()
		c.UpdatePids(procInfoArr)
		fmt.Printf("[PID SCAN] Total PIDs being monitored: %d\n", c.GetPidCount())
	}
}