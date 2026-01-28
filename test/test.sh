#!/bin/bash

WASM_FILE="sleep.wasm"
RUNTIMES=("wasmtime" "wasmedge")
PIDS=()

cleanup() {
    echo ""
    echo "--- Triggering cleanup ---"
    if [ ${#PIDS[@]} -eq 0 ]; then
        echo "No processes to clean up."
        return
    fi

    for pid in "${PIDS[@]}"; do
        if ps -p "$pid" > /dev/null; then
            echo "Terminating process with PID: $pid..."
            kill "$pid"
        else
            echo "Process with PID $pid no longer exists."
        fi
    done

    wait
    echo "--- Cleanup complete. All processes terminated. ---"
}

trap cleanup EXIT

if [ ! -f "$WASM_FILE" ]; then
    echo "Error: WASM file not found at '$WASM_FILE'"
    exit 1
fi

echo "Starting WASM runtimes..."
for runtime in "${RUNTIMES[@]}"; do
    if ! command -v "$runtime" &> /dev/null; then
        echo "Warning: Runtime '$runtime' not found. Skipping."
        continue
    fi

    echo "Launching: $runtime run $WASM_FILE"
    $runtime run "$WASM_FILE" &
    pid=$!
    PIDS+=($pid)
    echo " -> Started '$runtime' with PID: $pid"
done

if [ ${#PIDS[@]} -eq 0 ]; then
    echo "No runtimes were started. Exiting."
    exit 0
fi

echo ""
echo "Both processes are running in the background."
read -r -p "Press [Enter] to terminate them all... "