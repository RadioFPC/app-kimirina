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
	if value%50 ==