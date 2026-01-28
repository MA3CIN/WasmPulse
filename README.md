![WasmPulse Logo](media/WasmPulseLogo.png)
# WasmPulse
Analyzes resource usage of Server-side WebAssembly instances. Real-time CPU and Memory usage metrics in a Prometheus format, filling in the monitoring gap, acting as the cAdvisor for Wasm instances.

# Quickstart
To bring up just the WasmPulse container with necessary permissions, use:
```bash
docker compose up --no-build -d
```
To bring up WasmPulse, Prometheus and Grafana already preconfigured, use:
```bash
curl -sL https://raw.githubusercontent.com/MA3CIN/WasmPulse/feat/docker-compose-grafana-prometheus/release/docker-compose.yaml | docker-compose -f - up -d
```

# Release version
The release version supporst multiple Wasm instances simultaneously, and automatically finds all runnning services. It is also containerized and features an API for Prometheus to scrape the metrics.
For instructions, view the README.md in the [Release directory](release/README.md).

## Development
For development, use the command below to rebuild the image.
```bash
docker compose up --build -d
```


# Alpha version
The Alpha version represents the original IEEE CCNC 2026 paper - "WebAssembly as a Lightweight Path to Sustainable and High-Performance Cloud-Native Computing"  (https://repo.pw.edu.pl/info/article/WUT23a774b85f8f4a90b0557abaf2ab0702/, https://ccnc2026.ieee-ccnc.org/detailed-program). It is not recommended for Production use - instead, refer to the Release version.

It is based on shell scripts, which need filesystem permissions to run. For instructions, view the README.md in the [Alpha directory](alpha/README.md).