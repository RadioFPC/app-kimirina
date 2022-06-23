package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var roomManager *Manager

func main() {
	roomManager = NewRoomManager()
	router := gin.Default()
	router.SetHTMLTemplate(html)

	router.GET("/room/:roomid", roomGET)
	router.POST("/room/:roomid", roomPOST)
	router.DELETE("/room/:roomid", roomDELETE)
	router.GET("/stream/:roomid", stream)

	router.Run(":8080")
}

func stream(c *gin.Context) {
	roomid := c.Param("roomid")
	listene