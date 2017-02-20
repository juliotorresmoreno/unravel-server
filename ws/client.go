package ws

import (
	"encoding/json"
	"net/http"
	"time"

	"../models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	maxMessageSize = 8
)

//Client Conexiones websocket
type Client struct {
	conn    *websocket.Conn
	session *models.User
}

type user struct {
	session *models.User
	clients map[*Client]bool
	friends []string
}

//Clean Limpiar conexiones muertas
func (c user) Clean() {
	for key := range c.clients {
		err := key.conn.WriteMessage(websocket.PingMessage, make([]byte, 0))
		if err != nil {
			delete(c.clients, key)
		}
	}
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
	client.Listen()
}

//Listen es solo para disparar el evento close
func (c *Client) Listen() {
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}
	if _, ok := hub.clients[c.session.Usuario]; ok {
		delete(hub.clients[c.session.Usuario].clients, c)
		if len(hub.clients[c.session.Usuario].clients) == 0 {
			delete(hub.clients, c.session.Usuario)
			friends, _ := models.GetFriends(c.session.Usuario)
			estado, _ := json.Marshal(map[string]interface{}{
				"action":  "disconnect",
				"usuario": c.session.Usuario,
			})
			for _, e := range friends {
				hub.Send(e.Usuario, estado)
			}
		}
	}
	println("Conexion cerrada")
}
