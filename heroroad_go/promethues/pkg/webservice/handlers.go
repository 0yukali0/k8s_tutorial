package webservice

import (
	"net/http"

	"heroroad/pkg/metrics"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	metrics.GetPingMetrics().GetPingMetric().Count()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
