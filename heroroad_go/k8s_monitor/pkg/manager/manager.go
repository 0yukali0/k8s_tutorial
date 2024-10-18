package manager

import (
	"fmt"
	"context"
	"k8s_monitor/pkg/controller"
	"k8s_monitor/pkg/common"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

type Manager struct {
	Server      *gin.Engine
	Controllers []*controller.Controller
}

func (m *Manager) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	for _, ctrl := range m.Controllers {
		go ctrl.Run(ctx, 2)
	}
	defer cancel()
	m.Server.Run(fmt.Sprintf("127.0.0.1:%d", common.GetGlobalConfig().ServerPort))
}

func NewManager() *Manager {
	return &Manager{
		Server:      NewServer(),
		Controllers: NewControllers(),
	}
}

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/download", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func NewControllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewPodMonitor(),
	}
}
