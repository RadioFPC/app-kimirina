package main

import "html/template"

var html = template.Must(template.New("chat_room").Parse(`
<html> 
<head> 
    <title>{{.r