package chats

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"../../models"
	"../../ws"
	"../responses"
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
	e := make([]responses.Mensaje, len(q))
	if l > 0 {
		for i := 0; i < l; i++ {
			e[i] = responses.Mensaje{
				Action:          "mensaje",
				Usuario:   	 q[i].UsuarioEmisor,
				UsuarioReceptor: q[i].UsuarioReceptor,
				Mensaje:   	 q[i].Message,
				Fecha:     	 q[i].CreateAt,
			}
		}
	}
	resp, _ := json.Marshal(responses.SuccessData{
		Success: true,
		Data:    e,
	})
	w.Write(resp)
}

// Videollamada solicitud para la misma, debera ser aceptada por el usuario
func Videollamada(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	usuario := r.PostFormValue("usuario")
	resp, _ := json.Marshal(map[string]string{
		"action":          "videollamada",
		"usuario":         session.Usuario,
		"usuarioReceptor": usuario,
		"fecha":           strconv.Itoa(int(time.Now().Unix())),
	})
	hub.Send(session.Usuario, resp)
	hub.Send(usuario, resp)
	w.WriteHeader(http.StatusOK)
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
			resp, _ := json.Marshal(responses.Error{
				Success: false,
				Error:   "No puedes enviarte un mensaje a ti mismo",
			})
			w.Write(resp)
			return
		}
		chat := models.Chat{
			UsuarioEmisor:   session.Usuario,
			UsuarioReceptor: usuario,
			Message:         mensaje,
		}
		_, err := chat.Add()
		if err != nil {
			resp, _ := json.Marshal(responses.Error{
				Success: false,
				Error:   err.Error(),
			})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(responses.Success{
			Success: true,
			Message: "Enviado correctamente.",
		})
		w.Write(resp)
		resp, _ = json.Marshal(map[string]interface{} {
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
	resp, _ := json.Marshal(responses.Error{
		Success: false,
		Error:   "Not implemented",
	})
	w.Write(resp)
}
