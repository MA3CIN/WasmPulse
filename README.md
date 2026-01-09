# WasmPulse
WasmPulse - a Wasm performance evaluation platform. Filling in the monitoring gap for cAdvisor for Wasm instances.

# Quickstart
```bash
docker compose up --no-build -d
```

# Alpha version
The Alpha version represents the original IEEE CCNC 2026 paper - "WebAssembly as a Lightweight Path to Sustainable and High-Performance Cloud-Native Computing"  (https://ccnc2026.ieee-ccnc.org/detailed-program).

It is based on shell scripts, which need filesystem permissions to run. For instructions, view the README.md in the [Alpha directory](alpha/README.md).


# Release version
The release version supporst multiple Wasm instances simultaneously, and automatically finds all runnning services. It is also containerized and features an API.
For instructions, view the README.md in the [Release directory](release/README.md).

## Development
For development, use the command below to rebuild the image.
```bash
docker compose up --build -d
```
