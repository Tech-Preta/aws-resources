#!/bin/bash

# Lint script for aws-resources
# Runs various linting and formatting tools

set -e

echo "Running Go linters and formatters..."

# Format code
echo "Formatting code..."
go fmt ./...

# Vet code
echo "Vetting code..."
go vet ./...

# Check for golangci-lint
if command -v golangci-lint &> /dev/null; then
    echo "Running golangci-lint..."
    golangci-lint run
else
    echo "golangci-lint not found, skipping advanced linting"
    echo "Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
fi

# Check for goimports
if command -v goimports &> /dev/null; then
    echo "Running goimports..."
    goimports -w .
else
    echo "goimports not found, install with: go install golang.org/x/tools/cmd/goimports@latest"
fi

echo "Linting completed!"