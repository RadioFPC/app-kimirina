
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>