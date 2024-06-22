package metrics

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type PingMetrics struct {
	*PingMetric
}

func NewPingMetrics() *PingMetrics {
	return &PingMetrics{
		PingMetric: NewPingMetric(),
	}
}

func (p *PingMetrics) GetPingMetric() *PingMetric {
	return p.PingMetric
}

type PingMetric struct {
	PingCount prometheus.Counter
	lock      sync.RWMutex
}

func NewPingMetric() *PingMetric {
	s := &PingMetric{
		PingCount: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: Namespace,
				Subsystem: Ping,
				Name:      "ping_request_count",
				Help:      "Ping count",
			},
		),
		lock: sync.RWMutex{},
	}

	if err := prometheus.Register(s.PingCount); err != nil {
		fmt.Printf("%s", err.Error())
	}

	return s
}

func (p *PingMetric) Count() {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.PingCount.Inc()
}
