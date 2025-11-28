#!/bin/bash

# Build script for the backend server

set -e

echo "Building backend server..."

# Set build variables
VERSION=${VERSION:-"dev"}
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build the binary
go build -ldflags="-X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}' -X 'main.GitCommit=${GIT_COMMIT}'" \
    -o bin/server \
    ./cmd/server

echo "Build complete! Binary: bin/server"
echo "Version: ${VERSION}"
echo "Build Time: ${BUILD_TIME}"
echo "Git Commit: ${GIT_COMMIT}"
