package chats

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"../../helper"
	"../../models"
	"../../ws"
	"github.com/gorilla/mux"
)

var leido = models.Chat{Leido: 1}

// List obtiene la conversacion con el usuario solicitado
func List(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	var orm = models.GetXORM()
	var q = make([]models.Chat, 0)
	var usuario = vars["user"]
	var antesDe = r.URL.Query().Get("antesDe")
	var despuesDe = r.URL.Query().Get("despuesDe")

	var c = "(usuario_receptor = ? and usuario_emisor = ?) or (usuario_receptor = ? and usuario_emisor = ?)"
	if antesDe == "" && despuesDe == "" {
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario).Limit(10).OrderBy("id desc").Find(&q)
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario).Cols("leido").Update(leido)
	} else if antesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, antesDe)
		tiempo := tmp.String()[0:19]
		c = "(" + c + ") AND create_at < ?"
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario, tiempo).Limit(10).OrderBy("id desc").Find(&q)
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario, tiempo).Cols("leido").Update(leido)
	} else if despuesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, despuesDe)
		tiempo := tmp.String()[0:19]
		c = "(" + c + ") AND create_at > ?"
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario, tiempo).Limit(10).OrderBy("id desc").Find(&q)
		orm.Where(c, usuario, session.Usuario, session.Usuario, usuario, tiempo).Cols("leido").Update(leido)
	}
	l := len(q)
	e := make([]map[string]interface{}, len(q))
	if l > 0 {
		for i := 0; i < l; i++ {
			e[i] = map[string]interface{}{
				"action":          "mensaje",
				"usuario":         q[i].UsuarioEmisor,
				"usuarioReceptor": q[i].UsuarioReceptor,
				"mensaje":         q[i].Message,
				"fecha":           q[i].CreateAt,
			}
		}
	}
	resp, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    e,
	})
	w.Write(resp)
}

// Videollamada solicitud para la misma, debera ser aceptada por el usuario
func Videollamada(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	usuario := r.PostFormValue("usuario")
	resp, _ := json.Marshal(map[string]interface{}{
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
