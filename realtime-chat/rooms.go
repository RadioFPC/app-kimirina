
package main

import (
	"github.com/dustin/go-broadcast"
)

type Message struct {
	UserId string
	RoomId string
	Text   string
}

type Listener struct {
	RoomId string
	Chan   chan interface{}
}

type Manager struct {
	roomChannels map[string]broadcast.Broadcaster
	open         chan *Listener
	close        chan *Listener
	delete       chan string
	messages     chan *Message
}

func NewRoomManager() *Manager {
	manager := &Manager{
		roomChannels: make(map[string]broadcast.Broadcaster),
		open:         make(chan *Listener, 100),
		close:        make(chan *Listener, 100),
		delete:       make(chan string, 100),
		messages:     make(chan *Message, 100),
	}

	go manager.run()
	return manager
}

func (m *Manager) run() {
	for {
		select {
		case listener := <-m.open:
			m.register(listener)
		case listener := <-m.close:
			m.deregister(listener)
		case roomid := <-m.delete:
			m.deleteBroadcast(roomid)
		case message := <-m.messages:
			m.room(message.RoomId).Submit(message.UserId + ": " + message.Text)
		}
	}
}

func (m *Manager) register(listener *Listener) {
	m.room(listener.RoomId).Register(listener.Chan)
}

func (m *Manager) deregister(listener *Listener) {
	m.room(listener.RoomId).Unregister(listener.Chan)
	close(listener.Chan)
}

func (m *Manager) deleteBroadcast(roomid string) {
	b, ok := m.roomChannels[roomid]
	if ok {
		b.Close()
		delete(m.roomChannels, roomid)
	}
}

func (m *Manager) room(roomid string) broadcast.Broadcaster {