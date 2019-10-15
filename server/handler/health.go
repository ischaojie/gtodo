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
	Info   string `json:"info, omitempty"`
}

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {object} message "{"status":"OK", "info":""}"
// @Router /sd/health [get]
func HealthCheck(c *gin.Context) {
	mess := "ok"
	c.JSON(http.StatusOK, message{
		Status: mess,
	})
}

// * 检测硬盘使用量
// @Summary Checks the disk usage
// @Description Checks the disk usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {object} message "{"status":"OK", "info":"Free space: 44429MB (43GB) / 119674MB (116GB) | Used: 39%"}"
// @Router /sd/disk [get]
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
		Info:   mess,
	})

}

// * 检测cpu使用量
// @Summary Checks the cpu usage
// @Description Checks the cpu usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {object} message "{"status":"OK", "info":"Load average: 1.82, 1.05, 0.85 | Cores: 2"}"
// @Router /sd/cpu [get]
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
		Info:   mess,
	})
}

// * RAM 使用量
// @Summary Checks the ram usage
// @Description Checks the ram usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {object} message "{"status":"OK", "info":"Free space: 5293MB (5GB) / 7665MB (7GB) | Used: 69%"}"
// @Router /sd/ram [get]
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
		Info:   mess,
	})
}
