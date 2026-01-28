package discovery

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/shirou/gopsutil/v3/process"
)

func DiscoverWASM() []string{
	targets := []string{"wasmtime", "wasmedge", "wasmer", "spin", "wasmcloud", "wash"}
	pids := []string{}

	fmt.Println("Scanning for Wasm runtimes...")

	procs, err := process.Processes()
	if err != nil {
		fmt.Printf("Critical Error: Could not retrieve process list: %v\n", err)
		return pids
	}

	for _, p := range procs {
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()


		for _, target := range targets {
			if (strings.Contains(strings.ToLower(name), target) || 
				       strings.Contains(strings.ToLower(cmdline), target)){
						
						fmt.Printf("[FOUND] Runtime: %-10s | PID: %-6d | Command: %s\n", target, p.Pid, cmdline)
						pids = append(pids, strconv.Itoa(int(p.Pid)))

					   }
		}
	}
	return pids
}