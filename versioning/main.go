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

	// User only can be added by authorized person
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

	// version 2
	apiV2 := router.Group("/v2")

	apiV2.GET("use