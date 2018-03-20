package chats

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// VideoLlamada solicitud para la misma, debera ser aceptada por el usuario
func VideoLlamada(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)
	usuario := data.Get("usuario")
	resp, _ := json.Marshal(map[string]interface{}{
		"type":            "@chats/videocall",
		"action":          "videollamada",
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
