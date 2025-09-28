package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 首页，显示系统信息
	r.GET("/", func(c *gin.Context) {
		sysInfo := getSystemInfo()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"SysInfo": sysInfo,
		})
	})

	// 执行命令
	r.POST("/exec", func(c *gin.Context) {
		cmd := c.PostForm("command")
		out, err := execCommand(cmd)
		if err != nil {
			out += "\nError: " + err.Error()
		}
		c.JSON(http.StatusOK, gin.H{
			"output": out,
		})
	})

	r.Run(":8080") // 监听8080端口
}

// 获取系统信息
func getSystemInfo() map[string]interface{} {
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

// 执行命令
func execCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
