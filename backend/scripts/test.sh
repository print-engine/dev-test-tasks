#!/bin/bash

# Test script for the backend

set -e

echo "Running tests..."

# Run tests with coverage
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

echo ""
echo "Generating coverage report..."
go tool cover -func=coverage.out

echo ""
echo "Total coverage:"
go tool cover -func=coverage.out | grep total | awk '{print $3}'

# Optionally generate HTML coverage report
if [ "$1" == "--html" ]; then
    echo ""
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    echo "HTML report generated: coverage.html"
fi
