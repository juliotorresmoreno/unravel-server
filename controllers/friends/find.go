package friends

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unravel-server/models"
	"github.com/unravel-server/ws"
)

//ListFriends listado de amigos o personas con las que se puede chatear
func ListFriends(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var friends, _ = models.GetFriends(usuario)
	var estado, _ = json.Marshal(map[string]interface{}{
		"action":  "connect",
		"usuario": session.Usuario,
	})
	var amigos = make([]string, len(friends))
	var i = 0
	for el := range friends {
		friends[el].Conectado = hub.IsConnect(friends[el].Usuario)
		hub.Send(friends[el].Usuario, estado)
		amigos[i] = friends[el].Usuario
		i++
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    friends,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

// FindUser Busqueda de personas
func FindUser(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var query = r.URL.Query().Get("q")
	var usuario = r.URL.Query().Get("u")
	var users, _ = models.FindUser(session.Usuario, query, usuario)
	if len(users) == 1 && users[0].Relacion.EstadoRelacion == models.EstadoAceptado {
		users[0].Conectado = hub.IsConnect(users[0].Usuario)
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    users,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
