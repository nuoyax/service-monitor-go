package web

import (

	"net/http"
	"service-monitor/internal/executor"
	"service-monitor/internal/system"
	"runtime"
    "github.com/shirou/gopsutil/cpu"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 首页
	r.GET("/", func(c *gin.Context) {
		sysInfo := system.GetSystemInfo()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"SysInfo": sysInfo,
		})
	})

	// 多页面
	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page") + ".html"
		c.HTML(http.StatusOK, page, gin.H{})
	})

r.GET("/metrics", func(c *gin.Context) {
    vmStat := system.GetSystemInfo()
    cpuPercent, _ := cpu.Percent(0, false)
    upload, download := system.GetNetSpeed()
    wifi := system.GetWiFiStrength()

    c.JSON(http.StatusOK, gin.H{
        "os":           runtime.GOOS,
        "cpu_percent":  cpuPercent[0],        // float64
        "mem_percent":  vmStat["MemoryUsed"], // float64
        "disk_percent": vmStat["DiskUsed"],   // float64
        "wifi_percent": wifi,                  // float64
        "net_upload":   upload,                // float64
        "net_download": download,              // float64
    })
})




	// 执行命令
	r.POST("/exec", func(c *gin.Context) {
		cmd := c.PostForm("command")
		out, err := executor.ExecCommand(cmd)
		if err != nil {
			out += "\nError: " + err.Error()
		}
		c.JSON(http.StatusOK, gin.H{"output": out})
	})

	return r
}
