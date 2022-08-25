package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// version 1
	apiV1 := router.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

	// User only can be added by authoriz