package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	CYAN   = "\033[36m"
)

func bToGb(b uint64) float64 {
	return float64(b) / float64(1024*1024*1024)
}

func MemoryStatus() {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error Getting Memory Details %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	table.Append([]string{"Total", fmt.Sprintf("%.2f GB", bToGb(m.Total))})
	table.Append([]string{"Available", fmt.Sprintf("%.2f GB", bToGb(m.Available))})
	table.Append([]string{"Used", fmt.Sprintf("%.2f GB", bToGb(m.Used))})
	table.Append([]string{"Used Percent", fmt.Sprintf("%.2f%%", m.UsedPercent)})

	s, err := mem.SwapMemory()
	if err != nil {
		log.Fatalf("Error Getting Swap Details %v", err)
	}
	table.Append([]string{"Swap Total", fmt.Sprintf("%.2f GB", bToGb(s.Total))})
	table.Append([]string{"Swap Used", fmt.Sprintf("%.2f GB", bToGb(s.Used))})
	table.Append([]string{"Swap Free", fmt.Sprintf("%.2f GB", bToGb(s.Free))})

	fmt.Println(BLUE + "\nMemory Stats:" + RESET)
	table.Render()
}

func printDiskStats() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Error getting disk partitions:", err)
		return
	}

	fmt.Println(CYAN + "\nDisk Stats:" + RESET)
	for _, p := range parts {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Printf("Error getting disk usage for %s: %v\n", p.Mountpoint, err)
			continue
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Metric", "Value"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		)
		table.Append([]string{"Mountpoint", p.Mountpoint})
		table.Append([]string{"Total", fmt.Sprintf("%.2f GB", bToGb(usage.Total))})
		table.Append([]string{"Free", fmt.Sprintf("%.2f GB", bToGb(usage.Free))})
		table.Append([]string{"Used", fmt.Sprintf("%.2f GB", bToGb(usage.Used))})
		table.Append([]string{"Used Percent", fmt.Sprintf("%.2f%%", usage.UsedPercent)})
		table.Render()
	}
}

func printNetworkStats() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}

	fmt.Println(YELLOW + "\nNetwork Stats:" + RESET)
	for _, iface := range interfaces {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Metric", "Value"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		)
		table.Append([]string{"Interface", iface.Name})
		for _, addr := range iface.Addrs {
			table.Append([]string{"Address", addr.String()})
		}

		counters, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error getting network counters:", err)
			continue
		}
		for _, counter := range counters {
			if counter.Name == iface.Name {
				table.Append([]string{"Bytes Sent", fmt.Sprintf("%d B", counter.BytesSent)})
				table.Append([]string{"Bytes Recv", fmt.Sprintf("%d B", counter.BytesRecv)})
			}
		}
		table.Render()
	}
}

func printOSInfo() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	table.Append([]string{"GOOS", runtime.GOOS})
	table.Append([]string{"GOARCH", runtime.GOARCH})
	table.Append([]string{"CPUs", fmt.Sprintf("%d", runtime.NumCPU())})
	hostname, _ := os.Hostname()
	table.Append([]string{"Hostname", hostname})

	fmt.Println(GREEN + "\nOS Info:" + RESET)
	table.Render()
}

func printCPUStats() {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CPU", "Usage (%)"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	for i, p := range percent {
		table.Append([]string{fmt.Sprintf("CPU %d", i), fmt.Sprintf("%.2f%%", p)})
	}

	fmt.Println(RED + "\nCPU Stats:" + RESET)
	table.Render()
}

func printDiskIOStats() {
	ioCounters, err := disk.IOCounters()
	if err != nil {
		fmt.Println("Error getting disk I/O counters:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Device", "Read Count", "Write Count", "Read Bytes", "Write Bytes"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor},
	)
	for dev, io := range ioCounters {
		table.Append([]string{
			dev,
			fmt.Sprintf("%d", io.ReadCount),
			fmt.Sprintf("%d", io.WriteCount),
			humanize.Bytes(io.ReadBytes),
			humanize.Bytes(io.WriteBytes),
		})
	}

	fmt.Println(YELLOW + "\nDisk I/O Stats:" + RESET)
	table.Render()
}

func performNetworkSpeedTest() {
	fmt.Println(CYAN + "\nNetwork Speed Test:" + RESET)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("speedtest.exe", "--simple")
	case "darwin", "linux":
		cmd = exec.Command("speedtest-cli", "--simple")
	default:
		fmt.Println("Unsupported OS for network speed test.")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running speed test:", err)
		fmt.Println(string(output))
		return
	}

	outputStr := string(output)
	re := regexp.MustCompile(`Download: (.+) Mbps\nUpload: (.+) Mbps`)
	match := re.FindStringSubmatch(outputStr)

	if len(match) == 3 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Metric", "Value"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		)
		download, _ := strconv.ParseFloat(strings.TrimSpace(match[1]), 64)
		upload, _ := strconv.ParseFloat(strings.TrimSpace(match[2]), 64)
		table.Append([]string{"Download", fmt.Sprintf("%.2f Mbps", download)})
		table.Append([]string{"Upload", fmt.Sprintf("%.2f Mbps", upload)})
		table.Render()
	} else {
		fmt.Println("Could not parse speed test results. Output:", outputStr)
	}
}

func main() {
	memoryFlag := flag.Bool("memory", false, "Print memory stats")
	diskFlag := flag.Bool("disk", false, "Print disk stats")
	networkFlag := flag.Bool("network", false, "Print network stats")
	osFlag := flag.Bool("os", false, "Print OS info")
	cpuFlag := flag.Bool("cpu", false, "Print CPU stats")
	ioFlag := flag.Bool("io", false, "Print disk I/O stats")
	speedTestFlag := flag.Bool("speed", false, "Perform network speed test")
	allFlag := flag.Bool("all", false, "Print all stats")
	flag.Parse()

	if *allFlag {
		MemoryStatus()
		printDiskStats()
		printNetworkStats()
		printOSInfo()
		printCPUStats()
		printDiskIOStats()
		performNetworkSpeedTest()
	} else {
		if *memoryFlag {
			MemoryStatus()
		}
		if *diskFlag {
			printDiskStats()
		}
		if *networkFlag {
			printNetworkStats()
		}
		if *osFlag {
			printOSInfo()
		}
		if *cpuFlag {
			printCPUStats()
		}
		if *ioFlag {
			printDiskIOStats()
		}
		if *speedTestFlag {
			performNetworkSpeedTest()
		}
	}

	if !*memoryFlag && !*diskFlag && !*networkFlag && !*osFlag && !*cpuFlag && !*ioFlag && !*speedTestFlag && !*allFlag {
		fmt.Println("No flag specified. Use -h or --help for options.")
	}
}
