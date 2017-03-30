package users

import (
	"encoding/json"
	"net/http"

	"github.com/unravel-server/models"
	"github.com/unravel-server/ws"
)

// Find buscar usuarios
func Find(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var users, _ = models.FindUser(session.Usuario, r.URL.Query().Get("q"), r.URL.Query().Get("u"))
	if len(users) == 1 && users[0].Relacion != nil && users[0].Relacion.EstadoRelacion == models.EstadoAceptado {
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
