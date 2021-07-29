package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent"
)

const (
	// NewRelicTxnKey is the key used to retrieve the NewRelic Transaction from the context
	NewRelicTxnKey = "NewRelicTxnKey"
)

// NewRelicMonitoring is a middleware that starts a newrelic transaction, stores it in the context, then calls the next handl