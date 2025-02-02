# Makefile for Go project

# Variables
BINARY_NAME=system-stats
GO=go
GOFLAGS=-ldflags="-s -w"
BUILD_DIR=build
PLATFORMS=windows linux darwin
ARCHITECTURES=amd64

# Default target
all: build

# Build the project
build:
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

# Run the project
run:
	$(GO) run .

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Cross-compile for multiple platforms
cross-compile:
	@for GOOS in $(PLATFORMS); do \
		for GOARCH in $(ARCHITECTURES); do \
			export GOOS GOARCH; \
			$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-$$GOOS-$$GOARCH .; \
		done \
	done

# Install dependencies
deps:
	$(GO) mod download

# Format code
fmt:
	$(GO) fmt ./...

# Vet code
vet:
	$(GO) vet ./...

# Test code
test:
	$(GO) test ./...

# Lint code
lint:
	golint ./...

# Check code quality
check: fmt vet lint test

# Install the binary
install:
	$(GO) install .

# Uninstall the binary
uninstall:
	$(GO) clean -i

# Help message
help:
	@echo "Available targets:"
	@echo "  all           - Build the project (default)"
	@echo "  build         - Build the project"
	@echo "  run           - Run the project"
	@echo "  clean         - Clean build artifacts"
	@echo "  cross-compile - Cross-compile for multiple platforms"
	@echo "  deps          - Install dependencies"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  test          - Test code"
	@echo "  lint          - Lint code"
	@echo "  check         - Check code quality (fmt, vet, lint, test)"
	@echo "  install       - Install the binary"
	@echo "  uninstall     - Uninstall the binary"
	@echo "  help          - Show this help message"

.PHONY: all build run clean cross-compile deps fmt vet test lint check install uninstall help
