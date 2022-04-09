package main

import "github.com/dustin/go-broadcast"

var roomChannels = make(map[string]broadcast.Broadcaster)

func openListener(roomid string) chan interface{} {
	listener := make(chan interface{})
	room(roomid).Register(listener)
	return listener
}

fun