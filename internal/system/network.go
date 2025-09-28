package system

import (
	"time"

	"github.com/shirou/gopsutil/net"
)

var lastNetIO map[string]net.IOCountersStat
var lastCheckTime time.Time

// 网速
func GetNetSpeed() (upload float64, download float64) {
	ioStats, _ := net.IOCounters(true)
	now := time.Now()

	if lastNetIO == nil {
		lastNetIO = make(map[string]net.IOCountersStat)
		for _, stat := range ioStats {
			lastNetIO[stat.Name] = stat
		}
		lastCheckTime = now
		return 0, 0
	}

	duration := now.Sub(lastCheckTime).Seconds()
	var totalUpload, totalDownload uint64

	for _, stat := range ioStats {
		if last, ok := lastNetIO[stat.Name]; ok {
			uploadBytes := stat.BytesSent - last.BytesSent //:= 是一个非常常用的 短变量声明操作符 Go 会自动根据右边的值 推断变量类型。
			downloadBytes := stat.BytesRecv - last.BytesRecv
			totalUpload += uploadBytes
			totalDownload += downloadBytes
		}
		lastNetIO[stat.Name] = stat
	}

	lastCheckTime = now

	upload = float64(totalUpload*8) / duration / 1024 / 1024
	download = float64(totalDownload*8) / duration / 1024 / 1024
	return
}

// WiFi 强度（留接口）
func GetWiFiStrength() int {
	// TODO: 不同平台实现
	return 75
}
