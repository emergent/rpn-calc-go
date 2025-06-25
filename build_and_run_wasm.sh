#!/bin/bash

# Build the Go program to WebAssembly
echo "Building main.go to main.wasm..."
GOOS=js GOARCH=wasm go build -o main.wasm main.go

if [ $? -ne 0 ]; then
    echo "WebAssembly build failed!"
    exit 1
fi

# Copy wasm_exec.js if it doesn't exist or is outdated
GO_ROOT=$(go env GOROOT)
if [ ! -f wasm_exec.js ] || [ "$(find "$GO_ROOT/misc/wasm/wasm_exec.js" -prune -printf '%Y')" -gt "$(find "wasm_exec.js" -prune -printf '%Y')" ]; then
    echo "Copying wasm_exec.js..."
    cp "$GO_ROOT/misc/wasm/wasm_exec.js" .
else
    echo "wasm_exec.js is up to date."
fi

# Start a simple HTTP server
echo "Starting a simple HTTP server on port 8080..."
echo "Open your browser to http://localhost:8080"
python3 -m http.server 8080
