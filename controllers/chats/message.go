package chats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// Mensaje mensaje enviado por chat a los usuarios
func MensajeAdd(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
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
		id, err := chat.Add()
		if err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Enviado correctamente.",
		})
		resp, _ := json.Marshal(map[string]interface{}{
			"action":          "mensaje",
			"type":            "@chats/messagesAdd",
			"id":              id,
			"status":          "1",
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

// Mensaje mensaje enviado por chat a los usuarios
func MensajeEdit(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)
	id, _ := strconv.Atoi(data.Get("id"))
	status, _ := strconv.Atoi(data.Get("status"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	chat := models.Chat{}
	orm := db.GetXORM()
	defer orm.Close()
	_, err := orm.Where("id = ?", id).Get(&chat)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	usuario := session.Usuario
	if usuario == chat.UsuarioEmisor || usuario == chat.UsuarioReceptor {
		chat.Status = uint(status)
		if err := chat.Edit(id); err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Enviado correctamente.",
		})
		resp, _ := json.Marshal(map[string]interface{}{
			"action":          "mensaje",
			"type":            "@chats/messagesEdit",
			"id":              chat.Id,
			"status":          chat.Status,
			"usuario":         chat.UsuarioEmisor,
			"usuarioReceptor": chat.UsuarioReceptor,
			"mensaje":         chat.Message,
			"fecha":           time.Now(),
		})
		hub.Send(chat.UsuarioEmisor, resp)
		hub.Send(chat.UsuarioReceptor, resp)
		resp, _ = json.Marshal(map[string]interface{}{
			"action":          "mensaje",
			"type":            "@chats/videoCall",
			"usuario":         chat.UsuarioEmisor,
			"usuarioReceptor": chat.UsuarioReceptor,
			"videoCall":       chat.Status == 1 || chat.Status == 4,
		})
		hub.Send(chat.UsuarioEmisor, resp)
		hub.Send(chat.UsuarioReceptor, resp)
	}
}
