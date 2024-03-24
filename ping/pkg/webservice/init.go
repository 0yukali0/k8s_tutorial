package webservice

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once
var engine *gin.Engine

func GetGin() *gin.Engine {
	return engine
}

func init() {
	once.Do(func() {
		engine = gin.Default()
		for _, route := range webRoutes {
			switch route.Method {
			case "Get":
				engine.GET(route.Pattern, route.HandlerFunc)
			}
		}
	})
}
