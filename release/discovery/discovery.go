package discovery

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type WasmProcessInfo struct {
	PID         string
	RuntimeName string
	FileName    string
}

func init() {
	// Check if the host's /proc directory is mounted at /host/proc (for Docker/containers)
	if _, err := os.Stat("/host/proc"); err == nil {
		os.Setenv("HOST_PROC", "/host/proc")
		fmt.Println("[INIT] Detected '/host/proc' mount point.")
		fmt.Println("[INIT] Mode: Containerized (Host PID Monitoring Enabled)")
	} else {
		fmt.Println("[INIT] No '/host/proc' mount detected.")
		fmt.Println("[INIT] Mode: Standard (Bare Metal or Isolated Container)")
	}
}

func DiscoverWASM() []WasmProcessInfo {
	targets := []string{"wasmtime", "wasmedge", "wasmer", "spin", "wasmcloud", "wash"} // Todo - add more wasm runtimes!
	var foundProcesses []WasmProcessInfo

	procs, err := process.Processes()
	if err != nil {
		fmt.Printf("Critical Error: Could not retrieve process list: %v\n", err)
		return foundProcesses
	}

	for _, p := range procs {
		name, _ := p.Name()
		cmdSlice, _ := p.CmdlineSlice()
		cmdString := strings.Join(cmdSlice, " ")

		for _, target := range targets {
			if strings.Contains(strings.ToLower(name), target) ||
				strings.Contains(strings.ToLower(cmdString), target) {

				// Don't add duplicate PIDs to the array
				pidStr := strconv.Itoa(int(p.Pid))
				isDuplicate := false
				for _, existingPid := range foundProcesses {
					if existingPid.PID == pidStr {
						isDuplicate = true
						break
					}
				}

				if !isDuplicate {
					// Extract .wasm filename
					wasmFile := "unknown"
					for _, arg := range cmdSlice {
						if strings.HasSuffix(strings.ToLower(arg), ".wasm") {
							parts := strings.Split(arg, string(os.PathSeparator))
							wasmFile = parts[len(parts)-1]
							break
						}
					}

					runtimeDisplay := strings.ToUpper(target[:1]) + target[1:]

					fmt.Printf("[FOUND] Runtime: %-10s | PID: %-6s | File: %s\n", runtimeDisplay, pidStr, wasmFile)

					foundProcesses = append(foundProcesses, WasmProcessInfo{
						PID:         pidStr,
						RuntimeName: runtimeDisplay,
						FileName:    wasmFile,
					})
				}
			}
		}
	}
	return foundProcesses
}
