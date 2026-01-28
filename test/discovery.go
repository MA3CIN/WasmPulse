package main

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	targets := []string{"wasmtime", "wasmedge", "wasmer", "spin", "wasmcloud", "wash"}

	fmt.Println("Scanning for Wasm runtimes...")

	procs, err := process.Processes()
	if err != nil {
		fmt.Printf("Critical Error: Could not retrieve process list: %v\n", err)
		return
	}

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue 
		}

		cmdline, _ := p.Cmdline()

		for _, target := range targets {
			isMatch := strings.Contains(strings.ToLower(name), target) || 
				       strings.Contains(strings.ToLower(cmdline), target)

			if isMatch {
				fmt.Printf("[FOUND] Runtime: %-10s | PID: %-6d | Command: %s\n", target, p.Pid, cmdline)
				break
			}
		}
	}
}