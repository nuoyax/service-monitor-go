package system

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemInfo() map[string]interface{} {
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	cpuInfo, _ := cpu.Info()

	info := map[string]interface{}{
		"OS":         runtime.GOOS,
		"CPU":        cpuInfo[0].ModelName,
		"CPU_Cores":  runtime.NumCPU(),
		"Memory":     fmt.Sprintf("%.2f GB", float64(vmStat.Total)/1024/1024/1024),
		"MemoryUsed": fmt.Sprintf("%.2f%%", vmStat.UsedPercent),
		"Disk":       fmt.Sprintf("%.2f GB", float64(diskStat.Total)/1024/1024/1024),
		"DiskUsed":   fmt.Sprintf("%.2f%%", diskStat.UsedPercent),
	}
	return info
}
