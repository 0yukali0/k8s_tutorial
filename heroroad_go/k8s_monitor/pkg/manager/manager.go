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
	Events chan interface{}
}

func NewManager() *Manager {
	events := make(chan interface{}, 100)
	return &Manager{
		Server:      NewServer(),
		Controllers: NewControllers(events),
		Event: events,
	}
}

func (m *Manager) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	for _, ctrl := range m.Controllers {
		go ctrl.Run(ctx, 2)
	}
	defer cancel()
	go m.Server.Run(fmt.Sprintf("127.0.0.1:%d", common.GetGlobalConfig().ServerPort))
	for {
		select {
		// channel signal timer, ticker, chan, context
		case obj <- m.Events:
			GlobalPrint(obj)
		}
	}
}

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/pods", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, m.result)
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func NewControllers(events chan interface{}) []*controller.Controller {
	return []*controller.Controller{
		controller.NewPodMonitor(events),
	}
}

func (m *Manager) GlobalPrint(obj interface{}) {
	switch info := obj.(type) {
	case (*v1.Pod):
		m.result = NewPodInfo(info)
	case (*v1.K8sEvent):
	}
}

type PodInfo struct {
	Name string  `json:"name"`
	Namespace string `json:"namespace"`
	Image string `json:"image"`
	Phase string `json:"phase"`
}

func NewPodInfo(p *v1.Pod) {
	return PodInfo{
		Name: p.ObjectMeta.Name,
		Namespace: p.ObjectMeta.Namespace,
		Image: p.Spec.Containers[0].Image,
		Phase: p.Status.Phase,
	}
}