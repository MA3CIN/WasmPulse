package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/MA3CIN/WasmPulse/release/collector"
	"github.com/MA3CIN/WasmPulse/release/discovery"
)

func main() {
	pids := discovery.DiscoverWASM()
	fmt.Printf("Total runtimes found: %d\n", len(pids))

	collector.CollectMetrics(pids)
	serve()
}

func serve() {
		testMetric := promauto.NewCounter(prometheus.CounterOpts{
		Name: "test_metric",
		Help: "A simple metric that increments every 10 seconds",
	})

	go func() {
		for {
			testMetric.Inc()
			fmt.Println("Metric incremented at:", time.Now().Format("15:04:05"))
			time.Sleep(10 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server starting on :8080/metrics")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}