package main

import (
  "net/http"
  "time"
  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/mem"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/res", func(c *gin.Context) {
	percent, _ := cpu.Percent(time.Second, false)
	result := float64(0)
	for _, p := range percent {
		result += p/float64(len(percent))
	}
	memInfo, _ := mem.VirtualMemory()
    c.JSON(http.StatusOK, gin.H{
      "cpu": result,
	  "memory": memInfo.UsedPercent,
    })
  })
  r.StaticFile("/", "./res/dist/apps/fe/index.html")
  r.Static("/assets", "./res/dist/apps/fe/assets")
  r.Run()
}