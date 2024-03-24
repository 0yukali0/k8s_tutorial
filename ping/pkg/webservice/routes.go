package webservice

import (
	"github.com/gin-gonic/gin"
)

type route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type routes []route

var webRoutes = routes{
	route{
		Method:      "Get",
		Pattern:     "/ping",
		HandlerFunc: ping,
	},
}
