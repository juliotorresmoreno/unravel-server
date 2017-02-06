package ws

import "github.com/gorilla/websocket"

//Hub alacen de clientes websocket
type Hub struct {
	clients   map[string]*user
	broadcast chan []byte
}

var hub = &Hub{clients: make(map[string]*user)}

//GetHub devuelve el hub
func GetHub() *Hub {
	return hub
}

//IsConnect devuelve el estado de un usuario, conectado o desconectado.
func (hub Hub) IsConnect(user string) bool {
	usuario, ok := hub.clients[user]
	if ok && len(usuario.clients) > 0 {
		return true
	}
	return false
}

//Send enviar mensajes a los usuarios
func (hub Hub) Send(user string, mensaje []byte) {
	if client, ok := hub.clients[user]; ok {
		for conection := range client.clients {
			conection.conn.WriteMessage(websocket.TextMessage, mensaje)
		}
	}
}
