package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/examples/grpc/example1/gen/helloworld/v1"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Set 