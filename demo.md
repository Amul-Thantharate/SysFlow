# 🚀 SysFlow Demo Guide

## 📝 Features Showcase
- 💻 Real-time System Monitoring
  - Live CPU usage graphs
  - Memory utilization charts
  - Disk I/O metrics
  - Network traffic visualization
- 🔄 Advanced Monitoring Features
  - Per-core CPU tracking
  - Process resource usage
  - Network connection details
  - Disk I/O operations per second
- 🎨 Beautiful CLI Interface
  - Color-coded metrics
  - ASCII charts and graphs
  - Dynamic updates
  - Responsive layout

## 🖥️ Installation

### Linux
```bash
# Using make
make build-linux
./build/sysflow-linux

# Using go directly
GOOS=linux GOARCH=amd64 go build -o sysflow-linux
```

### macOS
```bash
# Using make
make build-macos
./build/sysflow-macos

# Using go directly
GOOS=darwin GOARCH=amd64 go build -o sysflow-macos
```

### Windows
```bash
# Using make
make build-windows
./build/sysflow-windows.exe

# Using go directly
GOOS=windows GOARCH=amd64 go build -o sysflow-windows.exe
```

## 🎯 Quick Start Examples

### Basic Monitoring
```bash
# Complete system overview
./sysflow all

# CPU monitoring only
./sysflow cpu --refresh=1s

# Memory monitoring
./sysflow memory --alert-threshold=90

# Disk usage
./sysflow disk --path=/home

# Network statistics
./sysflow network --interface=eth0
```

### Advanced Usage
```bash
# Monitor specific process
./sysflow process --name="nginx" --watch

# Export metrics to JSON
./sysflow all --export=json --output=metrics.json

# Remote monitoring
./sysflow --remote=192.168.1.100 --port=8080

# Custom refresh rate with alerts
./sysflow all --refresh=2s --alert-cpu=80 --alert-memory=90
```

## 🎨 Display Modes

### Minimal Mode
```bash
./sysflow all --minimal
```
```
CPU: 45% │█████████░░░░░░░░░░│ MEM: 6.2GB
DISK: 70% │███████████████░░░│ NET: 5MB/s
```

### Full Mode
```bash
./sysflow all --full
```
```
┌─ CPU Usage ───────────────┐ ┌─ Memory ──────────────┐
│ Core 0: 56% ████████░░░░░ │ │ Used:  8.5GB ███████░ │
│ Core 1: 32% ████░░░░░░░░░ │ │ Cache: 2.1GB ██░░░░░░ │
│ Core 2: 78% ██████████░░░ │ │ Free:  5.4GB ████░░░░ │
└─────────────────────────┘ └────────────────────────┘
```

### Process Mode
```bash
./sysflow process --top=5
```
```
PID    NAME      CPU   MEM   
1234   chrome    12.5% 2.1GB 
5678   node      8.2%  1.5GB 
9012   postgres  6.7%  1.2GB 
```

## 💡 Pro Tips & Tricks
1. 🔍 Use `-v` or `--verbose` for detailed output
   ```bash
   ./sysflow all -v
   ```

2. 🕒 Continuous monitoring with custom interval
   ```bash
   ./sysflow all --refresh=1s --duration=1h
   ```

3. 📊 Export data in different formats
   ```bash
   ./sysflow all --export=json
   ./sysflow all --export=csv
   ./sysflow all --export=prometheus
   ```

4. 🎯 Set custom alerts
   ```bash
   ./sysflow all --alert-cpu=80 --alert-memory=90 --alert-disk=95
   ```

5. 📱 Remote monitoring setup
   ```bash
   # On server
   ./sysflow server --port=8080

   # On client
   ./sysflow client --host=192.168.1.100 --port=8080
   ```

## 🎯 Common Use Cases

### 1. System Performance Debugging
```bash
./sysflow debug --full --log-level=debug
```

### 2. Resource Usage Monitoring
```bash
./sysflow monitor --process="nginx" --alert-cpu=80
```

### 3. Server Health Checks
```bash
./sysflow health --check-all --interval=5m
```

### 4. Performance Benchmarking
```bash
./sysflow benchmark --duration=1h --export=report.json
```

### 5. Continuous System Monitoring
```bash
./sysflow watch --critical-only --notify=slack
```

## 📊 Sample Output Formats

### JSON Output
```json
{
  "cpu": {
    "usage": 45.2,
    "cores": [
      {"id": 0, "usage": 56.7},
      {"id": 1, "usage": 32.1}
    ]
  },
  "memory": {
    "total": 16.0,
    "used": 8.5,
    "free": 7.5
  }
}
```

### CSV Output
```csv
timestamp,cpu_usage,memory_used,disk_usage,network_in
2025-01-30 13:32:51,45.2,8.5,70.1,5.2
```

## 🔧 Troubleshooting

### Common Issues
1. Permission denied
   ```bash
   sudo chmod +x ./sysflow-linux
   ```

2. Port already in use
   ```bash
   ./sysflow server --port=8081  # Use alternative port
   ```

3. High CPU usage
   ```bash
   ./sysflow all --minimal --refresh=5s  # Increase refresh interval
   ```
