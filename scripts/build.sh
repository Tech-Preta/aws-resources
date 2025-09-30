#!/bin/bash

# Build script for aws-resources
# This script builds the Go application with proper version information

set -e

# Get version information
VERSION=${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo "dev")}
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=${GIT_COMMIT:-$(git rev-parse HEAD 2>/dev/null || echo "unknown")}

# Build flags
LDFLAGS="-s -w"
LDFLAGS="$LDFLAGS -X main.version=$VERSION"
LDFLAGS="$LDFLAGS -X main.buildTime=$BUILD_TIME"
LDFLAGS="$LDFLAGS -X main.gitCommit=$GIT_COMMIT"

# Build directory
BUILD_DIR="bin"
BINARY_NAME="aws-resources"

echo "Building $BINARY_NAME version $VERSION..."
echo "Build time: $BUILD_TIME"
echo "Git commit: $GIT_COMMIT"

# Create build directory
mkdir -p $BUILD_DIR

# Build for current platform
go build -ldflags "$LDFLAGS" -o $BUILD_DIR/$BINARY_NAME ./cmd/aws-resources

echo "Build completed: $BUILD_DIR/$BINARY_NAME"