package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := gin.Default()
	r.GET("/test_stream", func(c *gin.Context) {
		w := c.Writer
		header := w.Header()
		header.Set("Transfer-En