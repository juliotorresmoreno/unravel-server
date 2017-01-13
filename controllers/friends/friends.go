package friends

import (
	"encoding/json"
	"net/http"

	"strings"

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
	var users = make([]models.User, 0)
	var relacion = make([]models.Relacion, 0)
	var orm = models.GetXORM()
	var str string
	if q := r.URL.Query().Get("q"); q != "" {
		w := strings.Split(q, " ")
		str = "usuario != ? AND (false"
		for _, v := range w {
			str = str + " OR (nombres LIKE '%" + v + "%' OR apellidos LIKE '%" + v + "%')"
		}
		str = str + ")"
	} else if u := r.URL.Query().Get("u"); u != "" {
		str = "Usuario != ? AND Usuario = '" + u + "'"
	} else {
		str = "Usuario != ?"
	}

	w.Header().Set("Content-Type", "application/json")
	if err := orm.Where(str, session.Usuario).Find(&users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta, _ := json.Marshal(responses.Error{Success: false, Error: err.Error()})
		w.Write(respuesta)
		return
	}

	str = "usuario_solicita = ? OR usuario_solicitado = ?"
	if err := orm.Where(str, session.Usuario, session.Usuario).Find(&relacion); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta, _ := json.Marshal(responses.Error{Success: false, Error: err.Error()})
		w.Write(respuesta)
		return
	}

	lengthUsers := len(users)
	lengthRelacion := len(relacion)
	list := make([]models.Friend, lengthUsers)
	for i := 0; i < lengthUsers; i++ {
		list[i] = models.Friend{
			Usuario:    users[i].Usuario,
			Nombres:    users[i].Nombres,
			Apellidos:  users[i].Apellidos,
			Estado:     "",
			Registrado: users[i].CreateAt,
		}
		for j := 0; j < lengthRelacion; j++ {
			solicita := relacion[j].UsuarioSolicita
			solicitado := relacion[j].UsuarioSolicitado
			if users[i].Usuario == solicita || users[i].Usuario == solicitado {
				if relacion[j].EstadoRelacion == models.EstadoSolicitado {
					list[i].Estado = "Solicitado"
				} else {
					list[i].Estado = "Amigos"
				}
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	respuesta, _ := json.Marshal(responses.ListFriends{
		Success: true,
		Data:    list,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

// Add agregar amigo
func Add(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = r.PostFormValue("user")
	if (usuario == "") {
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
	if len(relaciones) == 1 && relaciones[0].EstadoRelacion == 0 && relaciones[0].UsuarioSolicita == usuario {
		relaciones[0].EstadoRelacion = 1
		models.Update(relaciones[0].Id, relaciones[0])
		w.WriteHeader(http.StatusOK)
	} else if len(relaciones) == 0 {
		models.Add(models.Relacion{
			UsuarioSolicita: session.Usuario,
			UsuarioSolicitado: usuario,
			EstadoRelacion: 0,
		})
		w.WriteHeader(http.StatusOK)
	}
}
