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
	defer conn.Close()
	client := v1.NewGreeterClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")

		// Contact the server and print out its response.
		req := &v1.HelloRequest{Name: name}
		res, err := client.Sa