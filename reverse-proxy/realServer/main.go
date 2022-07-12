package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:2003"
)

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
		// in this handler, we just simply send some basic info back to proxy response.
		req := c.Request
		urlPath := fmt.Sprintf("http://%s%s", Addr, req.URL.Path)
		