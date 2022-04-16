package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func rateLimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(http.StatusServiceUnavailable, "you were automatically banned :)")
	}
}

func index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/room/hn")
}

func ro