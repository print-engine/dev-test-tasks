#!/bin/bash

# Development script with auto-reload

set -e

echo "Starting development server with auto-reload..."
echo "Install air if not already: go install github.com/air-verse/air@latest"
echo ""

# Check if air is installed
if ! command -v air &> /dev/null; then
    echo "air is not installed. Installing..."
    go install github.com/air-verse/air@latest
fi

# Run with air for hot reload
air -c .air.toml
