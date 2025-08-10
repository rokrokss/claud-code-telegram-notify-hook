#!/bin/bash

# Build notification binary for multiple platforms

echo "Building notification binaries..."

# # macOS (Intel and Apple Silicon)
# echo "Building for macOS (Intel)..."
# GOOS=darwin GOARCH=amd64 go build -o notification-darwin-amd64 notification.go

# echo "Building for macOS (Apple Silicon)..."
# GOOS=darwin GOARCH=arm64 go build -o notification-darwin-arm64 notification.go

# # Linux
# echo "Building for Linux (amd64)..."
# GOOS=linux GOARCH=amd64 go build -o notification-linux-amd64 notification.go

# echo "Building for Linux (arm64)..."
# GOOS=linux GOARCH=arm64 go build -o notification-linux-arm64 notification.go

# # Windows
# echo "Building for Windows (amd64)..."
# GOOS=windows GOARCH=amd64 go build -o notification-windows-amd64.exe notification.go

# Universal binary for current platform
echo "Building universal binary..."
go build -o notification-bin notification.go

echo "Build complete!"
ls -la notification-*