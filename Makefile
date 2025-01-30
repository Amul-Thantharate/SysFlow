.PHONY: all clean build-all build-linux build-macos build-windows

# Binary names
BINARY_LINUX=sysflow-linux
BINARY_MACOS=sysflow-macos
BINARY_WINDOWS=sysflow-windows.exe

# Build directory
BUILD_DIR=build

all: clean build-all

build-all: build-linux build-macos build-windows

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_LINUX)
	@echo "✅ Linux build complete"

build-macos:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_MACOS)
	@echo "✅ macOS build complete"

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_WINDOWS)
	@echo "✅ Windows build complete"

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"
