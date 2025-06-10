package system

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

// CPU-specific data
type CPUInfo struct {
	Usage []float64 // Per-core percentages
	Temp  float64   // Temperature (if available)
	Model string    // CPU model name
	Cores int       // Number of cores
}

// Memory-specific data
type MemInfo struct {
	Used    uint64  // Bytes used
	Total   uint64  // Total bytes
	Percent float64 // Percent used
	Free    uint64  // Free bytes
}

// Disk-specific data
type DiskInfo struct {
	Used    uint64  // Bytes used
	Total   uint64  // Total bytes
	Percent float64 // Percent used
	Free    uint64  // Free bytes
}

// GPU-specific data (limited in gopsutil)
type GPUInfo struct {
	Usage float64 // Percent usage
	Temp  float64 // Temperature
	Model string  // GPU model (if detectable)
}

func GetCPUInfo() (*CPUInfo, error) {
	ctx := context.Background()

	// CPU model info
	cpuInfo, err := cpu.InfoWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU info: %w", err)
	}

	// CPU usage per core
	cpuPercent, err := cpu.PercentWithContext(ctx, time.Second, true) // true = per CPU
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	model := "Unknown CPU"
	cores := len(cpuPercent)
	if len(cpuInfo) > 0 {
		model = cpuInfo[0].ModelName
		cores = int(cpuInfo[0].Cores)
	}

	return &CPUInfo{
		Usage: cpuPercent,
		Temp:  0.0, // TODO: Platform-specific temperature reading
		Model: model,
		Cores: cores,
	}, nil
}

func GetMemInfo() (*MemInfo, error) {
	ctx := context.Background()

	vmem, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %w", err)
	}

	return &MemInfo{
		Used:    vmem.Used,
		Total:   vmem.Total,
		Percent: vmem.UsedPercent,
		Free:    vmem.Available,
	}, nil
}

func GetDiskInfo() (*DiskInfo, error) {
	ctx := context.Background()

	// Use root partition - could be made configurable
	diskUsage, err := disk.UsageWithContext(ctx, "/") // Use "C:" on Windows
	if err != nil {
		return nil, fmt.Errorf("failed to get disk info: %w", err)
	}

	return &DiskInfo{
		Used:    diskUsage.Used,
		Total:   diskUsage.Total,
		Percent: diskUsage.UsedPercent,
		Free:    diskUsage.Free,
	}, nil
}

func GetGPUInfo() (*GPUInfo, error) {
	// gopsutil has very limited GPU support
	// This would need platform-specific implementation for real GPU data
	return &GPUInfo{
		Usage: 0.0,
		Temp:  0.0,
		Model: "GPU detection not implemented",
	}, nil
}

// Helper function to format bytes
func FormatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// Helper to get average CPU usage
func (c *CPUInfo) GetAvgUsage() float64 {
	if len(c.Usage) == 0 {
		return 0.0
	}

	total := 0.0
	for _, usage := range c.Usage {
		total += usage
	}
	return total / float64(len(c.Usage))
}
