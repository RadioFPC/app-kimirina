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
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer