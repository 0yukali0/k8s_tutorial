package main

import (
	"k8s_monitor/pkg/manager"
)

func main() {
	m := manager.NewManager()
	m.Run()
}
