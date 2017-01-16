package ws

import (
	"net/http"

	"../models"
	"github.com/gorilla/websocket"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	maxMessageSize = 8
)

type Client struct {
	conn    *websocket.Conn
	session *models.User
}

type user struct {
	session *models.User
	clients map[*Client]bool
}

// ServeWs aca es donde establenemos la conexion websocket con el usuario
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, session *models.User) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println(err)
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
	client.listen()
}


func (c *Client) listen() {
	defer func() {
		c.conn.Close()
		delete(hub.clients[c.session.Usuario].clients, c)
		println("Conexion cerrada")
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		if _, _, err := c.conn.ReadMessage(); err != nil {
			break
		}
		if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			break
		}
	}
}