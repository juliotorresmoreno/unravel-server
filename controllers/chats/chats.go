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

// List obtiene la conversacion con el usuario solicitado
func List(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	orm := models.GetXORM()
	q := make([]models.Chat, 0)
	usuario := vars["user"]

	c := "(usuario_receptor = ? and usuario_emisor = ?) or (usuario_receptor = ? and usuario_emisor = ?)"
	_ = orm.Where(c, usuario, session.Usuario, session.Usuario, usuario).Find(&q)
	l := len(q)
	e := make([]responses.Mensaje, len(q))
	if l > 0 {
		u := make([]models.User, 0)
		orm.Where("usuario = ?", usuario).Find(&u)

		for i := 0; i < l; i++ {
			if q[i].UsuarioEmisor != session.Usuario {
				e[i] = responses.Mensaje{
					Usuario:   u[0].Usuario,
					Nombres:   u[0].Nombres,
					Apellidos: u[0].Apellidos,
					Mensaje:   q[i].Message,
					Fecha:     q[i].CreateAt.Unix(),
				}
			} else {
				e[i] = responses.Mensaje{
					Usuario:   session.Usuario,
					Nombres:   session.Nombres,
					Apellidos: session.Apellidos,
					Mensaje:   q[i].Message,
					Fecha:     q[i].CreateAt.Unix(),
				}
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
		resp, _ = json.Marshal(map[string]string{
			"action":          "mensaje",
			"usuario":         session.Usuario,
			"usuarioReceptor": usuario,
			"mensaje":         mensaje,
			"fecha":           strconv.Itoa(int(time.Now().Unix())),
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
