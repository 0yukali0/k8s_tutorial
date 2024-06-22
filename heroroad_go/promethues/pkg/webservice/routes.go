package webservice

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type routes []route

var webRoutes = routes{
	route{
		Method:      "GET",
		Pattern:     "/ping",
		HandlerFunc: ping,
	},
	// Promethues metrics
	route{
		Method:  "GET",
		Pattern: "/metrics",
		HandlerFunc: func(c *gin.Context) {
			promhttp.Handler().ServeHTTP(c.Writer, c.Request)
		},
	},
	// pprof to retrive CPU, Memory profiling data.
	route{
		Method:      "GET",
		Pattern:     "/debug/pprof",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/heap",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/threadcreate",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/goroutine",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/allocs",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/block",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/mutex",
		HandlerFunc: gin.WrapF(pprof.Index),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/cmdline",
		HandlerFunc: gin.WrapF(pprof.Cmdline),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/profile",
		HandlerFunc: gin.WrapF(pprof.Profile),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/symbol",
		HandlerFunc: gin.WrapF(pprof.Symbol),
	},
	route{
		Method:      "GET",
		Pattern:     "debug/pprof/trace",
		HandlerFunc: gin.WrapF(pprof.Trace),
	},
}
