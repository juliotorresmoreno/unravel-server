package friends

import (
	"encoding/json"
	"net/http"

	"../../models"
	"../../ws"
	"../responses"
)

//ListFriends listado de amigos o personas con las que se puede chatear
func ListFriends(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var friends, _ = models.GetFriends(session.Usuario)
	respuesta, _ := json.Marshal(responses.ListFriends{
		Success: true,
		Data:    friends,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

// FindUser Busqueda de personas
func FindUser(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var users, _ = models.FindUser(session.Usuario, r.URL.Query().Get("q"), r.URL.Query().Get("u"))
	respuesta, _ := json.Marshal(responses.ListFriends{
		Success: true,
		Data:    users,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

// Add agregar amigo
func Add(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = r.PostFormValue("user")
	if (usuario == "" || usuario == session.Usuario) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var relaciones = make([]models.Relacion, 0)
	var orm = models.GetXORM()
	var str string = "(usuario_solicita = ? and usuario_solicitado = ?) or (usuario_solicita = ? and usuario_solicitado = ?)"
	if err := orm.Where(str, usuario, session.Usuario, session.Usuario, usuario).Find(&relaciones); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if len(relaciones) == 1 && relaciones[0].EstadoRelacion == 0 && relaciones[0].UsuarioSolicita == usuario {
		relaciones[0].EstadoRelacion = 1
		models.Update(relaciones[0].Id, relaciones[0])
		w.WriteHeader(http.StatusOK)
		respuesta, _ := json.Marshal(map[string]interface{} {
			"success": true,
			"relacion": relaciones[0],
			"estado": "Amigos",
		})
		w.Write(respuesta)
	} else if len(relaciones) == 1 {
		w.WriteHeader(http.StatusOK)
		respuesta, _ := json.Marshal(map[string]interface{} {
			"success": true,
			"relacion": relaciones[0],
			"estado": "Amigos",
		})
		w.Write(respuesta)
	} else if len(relaciones) == 0 {
		relacion := models.Relacion{
			UsuarioSolicita: session.Usuario,
			UsuarioSolicitado: usuario,
			EstadoRelacion: 0,
		}
		models.Add(relacion)
		w.WriteHeader(http.StatusOK)
		respuesta, _ := json.Marshal(map[string]interface{} {
			"success": true,
			"relacion": &relacion,
			"estado": "Solicitado",
		})
		w.Write(respuesta)
	}
}

// FindUser Busqueda de personas
func RejectFriend(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	models.RejectFriends(session.Usuario, r.PostFormValue("user"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}