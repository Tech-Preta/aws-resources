.PHONY: build test clean install help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=aws-resources
BINARY_PATH=bin/$(BINARY_NAME)

# Build the application
build:
	$(GOBUILD) -o $(BINARY_PATH) ./cmd/$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Install binary to GOPATH/bin
install: build
	cp $(BINARY_PATH) $(GOPATH)/bin/

# Development build with race detection
dev:
	$(GOBUILD) -race -o $(BINARY_PATH) ./cmd/$(BINARY_NAME)

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -cover ./...

# Format code
fmt:
	$(GOCMD) fmt ./...

# Vet code
vet:
	$(GOCMD) vet ./...

# Run all checks
check: fmt vet test

# Show help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install/update dependencies"
	@echo "  install       - Install binary to GOPATH/bin"
	@echo "  dev           - Build with race detection"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  check         - Run format, vet, and test"
	@echo "  help          - Show this help"