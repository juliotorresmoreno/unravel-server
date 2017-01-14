package users

import (
	"encoding/json"
	"net/http"

	"strings"

	"../../models"
	"../../ws"
	//"errors"
)


func findUsers(session *models.User, query string, id string) ([]models.User, error) {
	var users = make([]models.User, 0)
	var orm = models.GetXORM()
	var str string
	if id != "" {
		str = "Usuario != ? AND Usuario = '" + id + "'"
	} else if query != "" {
		w := strings.Split(query, " ")
		str = "usuario != ? AND (false"
		for _, v := range w {
			str = str + " OR (nombres LIKE '%" + v + "%' OR apellidos LIKE '%" + v + "%' OR usuario = '%" + v + "%')"
		}
		str = str + ")"
	} else {
		str = "Usuario != ?"
	}
	if err := orm.Where(str, session.Usuario).Find(&users); err != nil {
		return users, err //errors.New("Error desconocido")
	}
	return users, nil
}

func relacioUsers(session *models.User, users []models.User) ([]models.Friend, error) {
	var relacion = make([]models.Relacion, 0)
	var list = make([]models.Friend, 0)
	var orm = models.GetXORM()
	var str string
	str = "usuario_solicita = ? OR usuario_solicitado = ?"
	if err := orm.Where(str, session.Usuario, session.Usuario).Find(&relacion); err != nil {
		return list, err//errors.New("Error desconocido")
	}

	lengthUsers := len(users)
	lengthRelacion := len(relacion)
	list = make([]models.Friend, lengthUsers)
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
			} else {
				list[i].Estado = "Desconocido"
			}
		}
	}
	return list, nil
}


// Find busca los usuarios
func Find(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var list = make([]models.Friend, 0)
	var users = make([]models.User, 0)
	var err error
	if users, err = findUsers(session, r.URL.Query().Get("q"), r.URL.Query().Get("u")); err != nil {
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respuesta)
		return
	}
	if list, err = relacioUsers(session, users); err != nil {
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respuesta)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    list,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}