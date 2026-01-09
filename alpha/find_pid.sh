#!/bin/bash

# Find all PIDs with Wasm services ("Wasm" in name -> WasmTime, WasmEdge, etc.)
wasm_pids=$(ps -eo pid,comm | awk 'tolower($2) ~ /wasm/ {print $1}')

count=$(echo "$wasm_pids" | grep -c .)

if [ $count -eq 0 ]; then
    echo "No pids found, are you using a supported wasm runtime?"
else
    echo "$count pids found:"
    echo "$wasm_pids"
fi