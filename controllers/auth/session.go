package auth

import (
	"encoding/json"
	"net/http"

	"../../helper"
	"../../models"
	"../../ws"
)

// Session obtiene la session actual del usuario logueado
func Session(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	if session == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"success\": false}"))
		return
	}
	var _token string = helper.GetCookie(r, "token")
	if _token == "" {
		_token = r.URL.Query().Get("token")
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"session": map[string]string{
			"usuario":   session.Usuario,
			"nombres":   session.Nombres,
			"fullname":  session.FullName,
			"apellidos": session.Apellidos,
			"token":     _token,
		},
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
