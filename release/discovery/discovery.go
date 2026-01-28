package discovery

import (
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/shirou/gopsutil/v3/process"
)

func init() {
    // Check if the host's /proc directory is mounted at /host/proc (for Docker/containers)
    if _, err := os.Stat("/host/proc"); err == nil {
        // If it exists, set the environment variable that gopsutil uses
        os.Setenv("HOST_PROC", "/host/proc")
        fmt.Println("Host /proc detected. Running in Container on Host PID mode.")
    } else {
        fmt.Println("Running in Bare-Metal or Containerized PID mode.")
    }
}

func DiscoverWASM() []string {
    targets := []string{"wasmtime", "wasmedge", "wasmer", "spin", "wasmcloud", "wash"} // Todo - add more wasm runtimes!
    pids := []string{}

    procs, err := process.Processes()
    if err != nil {
        fmt.Printf("Critical Error: Could not retrieve process list: %v\n", err)
        return pids
    }

    for _, p := range procs {
        name, _ := p.Name()
        cmdline, _ := p.Cmdline()

        for _, target := range targets {
            if strings.Contains(strings.ToLower(name), target) ||
                strings.Contains(strings.ToLower(cmdline), target) {

                // Don't add duplicate PIDs to the array
                pidStr := strconv.Itoa(int(p.Pid))
                isDuplicate := false
                for _, existingPid := range pids {
                    if existingPid == pidStr {
                        isDuplicate = true
                        break
                    }
                }
                if !isDuplicate {
                    fmt.Printf("[FOUND] Runtime: %-10s | PID: %-6d | Command: %s\n", target, p.Pid, cmdline)
                    pids = append(pids, pidStr)
                }
            }
        }
    }
    return pids
}