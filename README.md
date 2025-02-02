# 🚀 SysFlow

A powerful cross-platform CLI tool for monitoring system resources and performance metrics in real-time. Built with Go, this tool provides comprehensive system insights with a beautiful and intuitive interface.

## ✨ Features
- 📊 Real-time system metrics monitoring
  - CPU core-wise utilization
  - Memory usage patterns
  - Disk I/O operations
  - Network throughput
  - Os and kernel information
  - SpeedTest for network
- 🚀 Lightweight and fast (< 1% CPU overhead)
- 🎨 Beautiful CLI interface with color-coded metrics
- 📈 Historical data tracking with customizable retention
- 🔔 Configurable alerts and notifications
- 📱 Remote monitoring capabilities
- 🔄 Auto-refresh with configurable intervals
- 📊 Export data in multiple formats (JSON, CSV, XML)

## 🎯 Key Benefits
- 🔍 Instant system performance insights
- ⚡ Quick problem identification
- 📈 Performance trend analysis
- 🚨 Proactive system monitoring
- 💾 Historical data analysis

## 🚀 Quick Start
```bash
# Clone the repository
git clone https://github.com/Amul-Thantharate/sysflow.git

# Build for your platform
make build-all    # Build for all platforms
make build-linux  # Build for Linux
make build-macos  # Build for macOS
make build-windows # Build for Windows

# Run the application
./build/sysflow-linux    # On Linux
./build/sysflow-macos    # On macOS
./build/sysflow-windows.exe  # On Windows
```

## 📦 Prerequisites
- Go 1.16 or higher
- Make (for building from source)
- Git
- Minimum system requirements:
  - RAM: 512MB
  - Storage: 100MB
  - CPU: 1 core

## 🛠️ Installation Options

### 📥 Using Go Install
```bash
go install github.com/Amul-Thantharate/sysflow@latest
```

### 🏗️ Building from Source
```bash
# Clone the repository
git clone https://github.com/Amul-Thantharate/sysflow.git
cd sysflow

# Build for all platforms
make build-all

# Or build for specific platform
make build-linux   # For Linux
make build-macos   # For macOS
make build-windows # For Windows
```

### 🐳 Using Docker
```bash
docker pull amulthantharate/sysflow
docker run -it --rm amulthantharate/sysflow
```

## 📖 Documentation
- [📚 Complete Documentation](docs/README.md)
- [🎮 Usage Guide](demo.md)
- [⚙️ Configuration Guide](docs/configuration.md)
- [🔧 Troubleshooting](docs/troubleshooting.md)

## 🛠️ Configuration
Configuration file is stored in `~/.config/sysflow/config.yaml`

### Sample Configuration
```yaml
refresh_rate: 5s
export_format: json
metrics:
  cpu: true
  memory: true
  disk: true
  network: true
alerts:
  cpu_threshold: 90
  memory_threshold: 85
```

## 📊 Supported Metrics
- CPU:
  - Overall usage percentage
  - Per-core utilization
  - Load average
  - Temperature (where available)
- Memory:
  - Total/Used/Available RAM
  - Swap usage
  - Page faults
- Disk:
  - Space utilization
  - I/O operations
  - Read/Write speeds
- Network:
  - Interface statistics
  - Bandwidth usage
  - Packet information
- Process:
  - Top processes by CPU/Memory
  - Process tree
  - Resource consumption

## 🎨 Color Coding
- 🟢 Green: Normal operation (0-60%)
- 🟡 Yellow: Warning level (61-85%)
- 🔴 Red: Critical level (86-100%)

## 🔧 Setting Up Aliases

To make SysFlow easier to use, you can set up aliases on your system. Here's how to do it on different operating systems:

### Linux

1. For Bash users (edit `~/.bashrc`):
```bash
# Add these lines to ~/.bashrc
alias sf='cd /path/to/SysFlow'
alias sysflow='(cd /path/to/SysFlow && go run main.go)'
```

2. For Zsh users (edit `~/.zshrc`):
```bash
# Add these lines to ~/.zshrc
alias sf='cd /path/to/SysFlow'
alias sysflow='(cd /path/to/SysFlow && go run main.go)'
```

After adding the aliases, run:
```bash
source ~/.bashrc  # for Bash
# OR
source ~/.zshrc   # for Zsh
```

### macOS

macOS uses the same approach as Linux. By default, it uses Zsh:

1. Edit `~/.zshrc`:
```bash
# Add these lines to ~/.zshrc
alias sf='cd /path/to/SysFlow'
alias sysflow='(cd /path/to/SysFlow && go run main.go)'
```

2. Apply changes:
```bash
source ~/.zshrc
```

### Windows

1. Using PowerShell (create/edit your PowerShell profile):
```powershell
# First, check if you have a profile:
echo $PROFILE

# If it doesn't exist, create it:
New-Item -Path $PROFILE -Type File -Force

# Add these lines to your profile:
function Set-SysFlowLocation { Set-Location C:\path\to\SysFlow }
Set-Alias -Name sf -Value Set-SysFlowLocation

function Start-SysFlow {
    Push-Location C:\path\to\SysFlow
    go run main.go
    Pop-Location
}
Set-Alias -Name sysflow -Value Start-SysFlow
```

2. Apply changes:
```powershell
. $PROFILE
```

Note: Replace `/path/to/SysFlow` or `C:\path\to\SysFlow` with your actual SysFlow installation path.

These aliases provide:
- `sf`: Quickly navigate to the SysFlow directory
- `sysflow`: Run SysFlow from any location

## 🤝 Contributing
We welcome contributions! Here's how you can help:

- 🐛 Report bugs
- 💡 Suggest features
- 📝 Improve documentation
- 🔧 Submit pull requests

Please read our [Contributing Guidelines](CONTRIBUTING.md) before getting started.

## 🧪 Testing
```bash
go test ./...
go test -race ./...  # Run tests with race detector
```

## 📈 Performance Impact
- CPU Usage: < 1%
- Memory Usage: < 50MB
- Disk I/O: Minimal (configurable logging)

## 🔒 Security
- 🔐 No root privileges required
- 🔒 Secure by default configuration
- 🛡️ Regular security updates

## 📝 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments
- Go community
- System monitoring tools that inspired this project
- All our contributors and users
