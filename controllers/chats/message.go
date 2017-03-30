package chats

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
	"github.com/unravel-server/ws"
)

// Mensaje mensaje enviado por chat a los usuarios
func Mensaje(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	usuario := r.PostFormValue("usuario")
	mensaje := r.PostFormValue("mensaje")
	tipo := r.PostFormValue("tipo")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if tipo == "usuario" {
		if usuario == session.Usuario {
			err := errors.New("No puedes enviarte un mensaje a ti mismo")
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		chat := models.Chat{
			UsuarioEmisor:   session.Usuario,
			UsuarioReceptor: usuario,
			Message:         mensaje,
		}
		_, err := chat.Add()
		if err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		resp, _ := json.Marshal(map[string]interface{}{
			"success": true,
			"message": "Enviado correctamente.",
		})
		w.Write(resp)
		resp, _ = json.Marshal(map[string]interface{}{
			"action":          "mensaje",
			"usuario":         session.Usuario,
			"usuarioReceptor": usuario,
			"mensaje":         mensaje,
			"fecha":           time.Now(),
		})
		hub.Send(session.Usuario, resp)
		hub.Send(usuario, resp)
		return
	}
	helper.DespacharError(w, errors.New("Not implemented"), http.StatusInternalServerError)
}
