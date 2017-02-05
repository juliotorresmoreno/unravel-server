package ws

import "github.com/gorilla/websocket"
import "time"

//Hub alacen de clientes websocket
type Hub struct {
	clients   map[string]*user
	broadcast chan []byte
}

var hub *Hub

func init() {
	hub = &Hub{clients: make(map[string]*user)}
}

//GetHub devuelve el hub
func GetHub() *Hub {
	return hub
}

//Run vigila y elimina las conexiones muertas
func Run() {
	for {
		for user, el := range hub.clients {
			el.Clean()
			if len(el.clients) == 0 {
				delete(hub.clients, user)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}

//IsConnect devuelve el estado de un usuario, conectado o desconectado.
func (hub Hub) IsConnect(user string) bool {
	usuario, ok := hub.clients[user]
	if !ok {
		return false
	}
	usuario.Clean()
	_, ok = hub.clients[user]
	if len(usuario.clients) == 0 {
		delete(hub.clients, user)
		return false
	}
	return ok
}

//Send enviar mensajes a los usuarios
func (hub Hub) Send(user string, mensaje []byte) {
	if client, ok := hub.clients[user]; ok {
		for conection := range client.clients {
			conection.conn.WriteMessage(websocket.TextMessage, mensaje)
		}
	}
}
