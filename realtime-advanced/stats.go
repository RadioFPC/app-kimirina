package main

import (
	"runtime"
	"sync"
	"time"

	"github.com/manucorporat/stats"
)

var (
	ips        = stats.New()
	messages   = stats.New()
	users      = stats.New()
	mutexStats sync.RWMutex
	savedStats map[string]uint64
)

func statsWorker() {
	c := time.Tick(1 * time.Second)
	var lastMallocs uint64
	var lastFrees uint64
	for range c {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)

		mutexStats.Lock