#!/bin/bash

# Test script for aws-resources
# Runs comprehensive tests with coverage

set -e

echo "Running Go tests..."

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# Generate coverage report
if command -v go &> /dev/null; then
    echo "Generating coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    echo "Coverage report generated: coverage.html"
fi

# Display coverage summary
go tool cover -func=coverage.out | tail -1

echo "Tests completed successfully!"