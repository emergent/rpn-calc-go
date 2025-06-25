#!/bin/bash

# Build the Go program to WebAssembly
echo "Building main.go to main.wasm..."
GOOS=js GOARCH=wasm go build -o main.wasm main.go

if [ $? -ne 0 ]; then
    echo "WebAssembly build failed!"
    exit 1
fi

# Copy wasm_exec.js if it doesn't exist
GO_ROOT=$(go env GOROOT)
if [ ! -f wasm_exec.js ]; then
    echo "Copying wasm_exec.js..."
    cp "$GO_ROOT/misc/wasm/wasm_exec.js" .
else
    echo "wasm_exec.js already exists."
fi

# Start a simple HTTP server
echo "Starting a simple HTTP server on port 8080..."
echo "Open your browser to http://localhost:8080"
python3 -m http.server 8080 || python -m SimpleHTTPServer 8080 # Try python3, then python2

