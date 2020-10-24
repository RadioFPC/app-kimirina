package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define