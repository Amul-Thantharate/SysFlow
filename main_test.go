package main

import (
	"testing"
)

// TestMemoryStatus tests the MemoryStatus function
func TestMemoryStatus(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	MemoryStatus()
}

// TestPrintDiskStats tests the printDiskStats function
func TestPrintDiskStats(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	printDiskStats()
}

// TestPrintNetworkStats tests the printNetworkStats function
func TestPrintNetworkStats(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	printNetworkStats()
}

// TestPrintOSInfo tests the printOSInfo function
func TestPrintOSInfo(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	printOSInfo()
}

// TestPrintCPUStats tests the printCPUStats function
func TestPrintCPUStats(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	printCPUStats()
}

// TestPrintDiskIOStats tests the printDiskIOStats function
func TestPrintDiskIOStats(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	printDiskIOStats()
}

// TestPerformNetworkSpeedTest tests the performNetworkSpeedTest function
func TestPerformNetworkSpeedTest(t *testing.T) {
	// This is a basic test to ensure the function runs without errors
	performNetworkSpeedTest()
}

// TestBToGb tests the bToGb function
func TestBToGb(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected float64
	}{
		{"1GB", 1024 * 1024 * 1024, 1.0},
		{"2GB", 2 * 1024 * 1024 * 1024, 2.0},
		{"0GB", 0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bToGb(tt.input)
			if result != tt.expected {
				t.Errorf("bToGb(%d) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
