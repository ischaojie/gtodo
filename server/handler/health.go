package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type message struct {
	Status string `json:"status"`
	 Info string `json:"info, omitempty"`
}

func HealthCheck(c *gin.Context) {
	mess := "ok"
	c.JSON(http.StatusOK, message{
		Status: mess,
	})
}

// * 检测硬盘使用量
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	mess := fmt.Sprintf("Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.JSON(status, message{
		Status: text,
		Info:  mess,
	})

}

// * 检测cpu使用量
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	mess := fmt.Sprintf("Load average: %.2f, %.2f, %.2f | Cores: %d", l1, l5, l15, cores)
	c.JSON(status, message{
		Status: text,
		Info:  mess,
	})
}

// * RAM 使用量
func RAMCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	mess := fmt.Sprintf("Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.JSON(status, message{
		Status: text,
		Info:  mess,
	})
}

