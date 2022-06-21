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
	roomManag