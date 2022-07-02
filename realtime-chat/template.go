package main

import "html/template"

var html = template.Must(template.New("chat_room").Parse(`
<html> 
<head> 
    <title>{{.roomid}}</title>
    <link rel="stylesheet" type="text/css" href="http://meyerweb.com/eric/tools/css/reset/reset.css">
    <script src="http://ajax.go