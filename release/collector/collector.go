package collector

import (
	"log"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/process"
)

type PidCollector struct {
	rssGauge *prometheus.GaugeVec
	cpuGauge *prometheus.GaugeVec
	pids     []string
	mu       sync.RWMutex // Mutex to protect the pids slice
}

func NewPidCollector() *PidCollector {
	return &PidCollector{
		rssGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "process_resident_memory_bytes",
				Help: "Resident Set Size of a process.",
			},
			[]string{"pid"},
		),
		cpuGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "process_cpu_usage_percent",
				Help: "CPU usage of a process.",
			},
			[]string{"pid"},
		),
		pids: []string{},
	}
}

func (c *PidCollector) UpdatePids(newPids []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	existingPids := make(map[string]bool)
	for _, pid := range c.pids {
		existingPids[pid] = true
	}

	for _, pid := range newPids {
		if !existingPids[pid] {
			log.Printf("New PID discovered: %s. Adding to monitor list.", pid)
			c.pids = append(c.pids, pid)
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

    var activePids []string

    for _, pidStr := range c.pids {
        pid, err := strconv.ParseInt(pidStr, 10, 32)
        if err != nil {
            log.Printf("Error parsing PID '%s': %v", pidStr, err)
            continue
        }

        proc, err := process.NewProcess(int32(pid))
        
		if err != nil {
            log.Printf("Process %s terminated. Removing from metrics.", pidStr)
            c.rssGauge.Delete(prometheus.Labels{"pid": pidStr})
            c.cpuGauge.Delete(prometheus.Labels{"pid": pidStr})
            continue
        }

		activePids = append(activePids, pidStr)

        if memInfo, err := proc.MemoryInfo(); err == nil {
            c.rssGauge.With(prometheus.Labels{"pid": pidStr}).Set(float64(memInfo.RSS))
        }

        if cpuPercent, err := proc.CPUPercent(); err == nil {
            c.cpuGauge.With(prometheus.Labels{"pid": pidStr}).Set(cpuPercent)
        }
    }

    c.pids = activePids

    c.rssGauge.Collect(ch)
    c.cpuGauge.Collect(ch)
}