package metrics

import (
	"sync"
)

const (
	Namespace              = "HeroRoad"
	Ping                   = "ping"
	InvalidByteReplacement = '_'
)

var once sync.Once
var m *Metrics

type Metrics struct {
	*PingMetrics
	lock sync.RWMutex
}

func init() {
	once.Do(func() {
		m = &Metrics{
			PingMetrics: NewPingMetrics(),
			lock:        sync.RWMutex{},
		}
	})
}

func GetMetrics() *Metrics {
	return m
}

func GetPingMetrics() *PingMetrics {
	return m.PingMetrics
}
