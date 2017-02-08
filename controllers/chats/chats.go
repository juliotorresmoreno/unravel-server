package chats

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"../../helper"
	"../../models"
	"../../ws"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

var leido = models.Chat{Leido: 1}

// GetConversacion obtiene la conversacion con el usuario solicitado
func GetConversacion(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	var orm = models.GetXORM()
	var resultado = make([]models.Chat, 0)
	var usuario = vars["user"]
	var antesDe = r.URL.Query().Get("antesDe")
	var despuesDe = r.URL.Query().Get("despuesDe")

	var cond = "(usuario_receptor = ? and usuario_emisor = ?) or (usuario_receptor = ? and usuario_emisor = ?)"

	var updater *xorm.Session
	var consultor *xorm.Session
	if antesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, antesDe)
		tiempo := tmp.String()[0:19]
		cond = "(" + cond + ") AND create_at < ?"
		updater = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario, tiempo)
		consultor = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario, tiempo)
	} else if despuesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, despuesDe)
		tiempo := tmp.String()[0:19]
		cond = "(" + cond + ") AND create_at > ?"
		updater = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario, tiempo)
		consultor = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario, tiempo)
	} else {
		updater = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario)
		consultor = orm.Where(cond, usuario, session.Usuario, session.Usuario, usuario)
	}
	updater.Cols("leido").Update(leido)
	consultor.OrderBy("id desc").Limit(10).Find(&resultado)
	length := len(resultado)
	conversacion := make([]map[string]interface{}, length)
	if length > 0 {
		for i := 0; i < length; i++ {
			conversacion[i] = map[string]interface{}{
				"action":          "mensaje",
				"usuario":         resultado[i].UsuarioEmisor,
				"usuarioReceptor": resultado[i].UsuarioReceptor,
				"mensaje":         resultado[i].Message,
				"fecha":           resultado[i].CreateAt,
			}
		}
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    conversacion,
	})
	w.Write(respuesta)
}

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
