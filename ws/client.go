package ws

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"../models"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn    *websocket.Conn
	session *models.User
}

type user struct {
	session *models.User
	clients map[*Client]bool
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, session *models.User) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{conn: conn, session: session}
	if _, ok := hub.clients[session.Usuario]; ok == false {
		hub.clients[session.Usuario] = &user{
			session: session,
			clients: make(map[*Client]bool),
		}
	}
	hub.clients[session.Usuario].clients[client] = true
}