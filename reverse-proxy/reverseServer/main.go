
package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	// this is our reverse server ip address
	ReverseServerAddr = "127.0.0.1:2002"
)

// maybe we can have many real server addresses and do some load balanced strategy.
var RealAddr = []string{
	"http://127.0.0.1:2003",
}

// a fake function that we can do strategy here.
func getLoadBalanceAddr() string {