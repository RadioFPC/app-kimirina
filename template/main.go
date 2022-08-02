package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDat