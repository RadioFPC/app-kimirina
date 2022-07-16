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
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<html>
					<body>
		`))
		w.(http.Flusher).