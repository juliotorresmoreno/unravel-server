package ws

import (
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients   map[string]*user
	broadcast chan []byte
}

var hub *Hub

func init() {
	hub = &Hub{clients: make(map[string]*user)}
}

func GetHub() *Hub {
	return hub
}

func (hub Hub) Send(user string, mensaje []byte) {
	if client, ok := hub.clients[user]; ok {
		for conection := range client.clients {
			conection.conn.WriteMessage(websocket.TextMessage, mensaje)
		}
	}
}
