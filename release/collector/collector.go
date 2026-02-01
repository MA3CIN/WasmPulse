package collector

import (
	"log"
	"strconv"
	"sync"

	"github.com/MA3CIN/WasmPulse/release/discovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/process"
)

type PidCollector struct {
	rssGauge *prometheus.GaugeVec
	cpuGauge *prometheus.GaugeVec
	pids     []discovery.WasmProcessInfo
	mu       sync.RWMutex // Mutex to protect the pids slice
}

func NewPidCollector() *PidCollector {
	prometheus_labels := []string{"wasm_file", "runtime", "pid"}

	return &PidCollector{
		rssGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "process_resident_memory_bytes",
				Help: "Resident Set Size of a process.",
			},
			prometheus_labels,
		),
		cpuGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "process_cpu_usage_percent",
				Help: "CPU usage of a process.",
			},
			prometheus_labels,
		),
		pids: []discovery.WasmProcessInfo{},
	}
}

func (c *PidCollector) UpdatePids(newProcs []discovery.WasmProcessInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	existingPids := make(map[string]bool)
	for _, proc := range c.pids {
		existingPids[proc.PID] = true
	}

	for _, proc := range newProcs {
		if !existingPids[proc.PID] {
			log.Printf("New PID discovered: %s (%s - %s). Adding to monitor list.", proc.PID, proc.FileName, proc.RuntimeName)
			c.pids = append(c.pids, proc)
		}
	}
}

func (c *PidCollector) GetPidCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.pids)
}

func (c *PidCollector) Describe(ch chan<- *prometheus.Desc) {
	c.rssGauge.Describe(ch)
	c.cpuGauge.Describe(ch)
}

func (c *PidCollector) Collect(ch chan<- prometheus.Metric) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var activeProcs []discovery.WasmProcessInfo

	for _, procInfo := range c.pids {
		pidInt, err := strconv.ParseInt(procInfo.PID, 10, 32)
		if err != nil {
			log.Printf("Error parsing PID '%s': %v", pidInt, err)
			continue
		}

		proc, err := process.NewProcess(int32(pidInt))

		labels := prometheus.Labels{
			"wasm_file": procInfo.FileName,
			"runtime":   procInfo.RuntimeName,
			"pid":       procInfo.PID,
		}

		if err != nil {
			log.Printf("Process %s terminated. Removing from metrics.", procInfo.PID)
			c.rssGauge.Delete(labels)
			c.cpuGauge.Delete(labels)
			continue
		}

		activeProcs = append(activeProcs, procInfo)

		if memInfo, err := proc.MemoryInfo(); err == nil {
			c.rssGauge.With(labels).Set(float64(memInfo.RSS))
		}

		if cpuPercent, err := proc.CPUPercent(); err == nil {
			c.cpuGauge.With(labels).Set(cpuPercent)
		}
	}

	c.pids = activeProcs

	c.rssGauge.Collect(ch)
	c.cpuGauge.Collect(ch)
}
