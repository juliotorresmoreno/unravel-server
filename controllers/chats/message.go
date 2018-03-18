package chats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// Mensaje mensaje enviado por chat a los usuarios
func Mensaje(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)
	usuario := data.Get("usuario")
	mensaje := data.Get("mensaje")
	tipo := data.Get("tipo")
	w.Header().Set("Content-Type", "application/json")

	if tipo == "usuario" {
		w.WriteHeader(http.StatusOK)
		if usuario == session.Usuario {
			err := fmt.Errorf("No puedes enviarte un mensaje a ti mismo")
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
			"type":            "@chats/messagesAdd",
			"usuario":         session.Usuario,
			"usuarioReceptor": usuario,
			"mensaje":         mensaje,
			"fecha":           time.Now(),
		})
		hub.Send(session.Usuario, resp)
		hub.Send(usuario, resp)
		return
	}
	helper.DespacharError(w, fmt.Errorf("Not implemented"), http.StatusInternalServerError)
}
