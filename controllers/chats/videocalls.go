package chats

import (
	"encoding/json"
	"net/http"
	"time"

	"../../models"
	"../../ws"
)

// VideoLlamada solicitud para la misma, debera ser aceptada por el usuario
func VideoLlamada(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = r.PostFormValue("usuario")
	var tipo = r.PostFormValue("tipo")
	resp, _ := json.Marshal(map[string]interface{}{
		"action":          "videollamada",
		"usuario":         session.Usuario,
		"tipo":            tipo,
		"usuarioReceptor": usuario,
		"fecha":           time.Now(),
	})
	hub.Send(session.Usuario, resp)
	hub.Send(usuario, resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}

// RechazarVideoLlamada solicitud para la misma, debera ser aceptada por el usuario
func RechazarVideoLlamada(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	usuario := r.PostFormValue("usuario")
	resp, _ := json.Marshal(map[string]interface{}{
		"action":          "rechazarvideollamada",
		"usuario":         session.Usuario,
		"usuarioReceptor": usuario,
		"fecha":           time.Now(),
	})
	hub.Send(session.Usuario, resp)
	hub.Send(usuario, resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}
