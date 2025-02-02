package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
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

// MemoryStatus prints memory usage statistics
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

// PrintDiskStats prints disk usage statistics
func PrintDiskStats() {
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

// PrintNetworkStats prints network interface statistics
func PrintNetworkStats() {
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

// PrintOSInfo prints OS information
func PrintOSInfo() {
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

// PrintCPUStats prints CPU usage statistics
func PrintCPUStats() {
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

// PrintUsers prints a list of users and their information
func PrintUsers() {
	users, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	table.Append([]string{"Username", users.Username})
	table.Append([]string{"Name", users.Name})
	table.Append([]string{"Home Directory", users.HomeDir})
	table.Append([]string{"User ID", users.Uid})
	table.Append([]string{"Group ID", users.Gid})

	fmt.Println(CYAN + "\nUser Info:" + RESET)
	table.Render()
}

func PrintOpenPorts() {
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("Error getting network interfaces:", err)
        return
    }

    fmt.Println(BLUE + "\nOpen Ports and IP Addresses:" + RESET)

    for _, iface := range interfaces {
        addrs := iface.Addrs
        for _, addr := range addrs {
            fmt.Printf("Interface: %s, Address: %s\n", iface.Name, addr.Addr)
        }
    }
}

func main() {
	memoryFlag := flag.Bool("memory", false, "Print memory stats")
	diskFlag := flag.Bool("disk", false, "Print disk stats")
	networkFlag := flag.Bool("network", false, "Print network stats")
	osFlag := flag.Bool("os", false, "Print OS info")
	cpuFlag := flag.Bool("cpu", false, "Print CPU stats")
	tempFlag := flag.Bool("temp", false, "Print system temperature")
	usersFlag := flag.Bool("users", false, "Print user info")
	openPorts := flag.Bool("ports", false, "Print all open ports")
	allFlag := flag.Bool("all", false, "Print all stats")
	flag.Parse()

	if *allFlag {
		MemoryStatus()
		PrintDiskStats()
		PrintNetworkStats()
		PrintOSInfo()
		PrintCPUStats()
		PrintUsers()
		PrintOpenPorts()
	} else {
		if *memoryFlag {
			MemoryStatus()
		}
		if *diskFlag {
			PrintDiskStats()
		}
		if *networkFlag {
			PrintNetworkStats()
		}
		if *osFlag {
			PrintOSInfo()
		}
		if *cpuFlag {
			PrintCPUStats()
		}
		if *usersFlag {
			PrintUsers()
		}
		if *openPorts {
			PrintOpenPorts()
		}
	}

	if !*memoryFlag && !*diskFlag && !*networkFlag && !*osFlag && !*cpuFlag && !*tempFlag && !*usersFlag && !*openPorts && !*allFlag {
		fmt.Println("No flag specified. Use -h or --help for options.")
	}
}
