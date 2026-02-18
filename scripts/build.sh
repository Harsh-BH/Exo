#!/bin/bash

# Define app name
APP_NAME="exo"
OUTPUT_DIR="bin"

# Create output directory
mkdir -p $OUTPUT_DIR

echo "Building EXO CLI..."

# Linux (amd64)
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-linux-amd64 main.go

# macOS (amd64 - Intel)
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-darwin-amd64 main.go

# macOS (arm64 - Apple Silicon)
echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o $OUTPUT_DIR/$APP_NAME-darwin-arm64 main.go

# Windows (amd64)
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-windows-amd64.exe main.go

echo "Build complete! Binaries are in $OUTPUT_DIR/"
ls -lh $OUTPUT_DIR/
