package main

import (
	"gorm.io/gorm"
	"runtime"
)

type Memstat struct {
	total uint64
	used uint64
}

func NewMemstat() *Memstat {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return &Memstat{
		total: mem.TotalAlloc,
		used: mem.Mallocs
	}
}

type GoStatus struct {
	gorm.Model
	Mem: *Memstat `gorm: "embedded"`
	CPU: int
	Goroutine int
}

func NewGoStatus() *GoStatus {
	return &GoStatus{
		CPU: runtime.NumCPU(),
		Goroutine: runtime.NumGoroutine(),
		Mem: NewMemstat()
	}
}
