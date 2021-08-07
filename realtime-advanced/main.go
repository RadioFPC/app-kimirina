package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runt