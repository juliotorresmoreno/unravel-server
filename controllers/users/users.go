package users

import (
	"encoding/json"
	"net/http"

	"../../models"
	"../../ws"
)

// Find buscar usuarios
func Find(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var users, _ = models.FindUser(session.Usuario, r.URL.Query().Get("q"), r.URL.Query().Get("u"))
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    users,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
