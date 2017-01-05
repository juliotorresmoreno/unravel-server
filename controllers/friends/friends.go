package friends

import (
	"encoding/json"
	"net/http"

	"strings"

	"../../models"
	"../../ws"
	"../responses"
)

// ListFriends listado de amigos o personas con las que se puede chatear
func ListFriends(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var users = make([]models.User, 0)
	var orm = models.GetXORM()
	var err error
	var str string
	if q := r.URL.Query().Get("q"); q != "" {
		w := strings.Split(q, " ")
		str = "usuario != ? AND (false"
		for _, v := range w {
			str = str + " OR (nombres LIKE '%" + v + "%' OR apellidos LIKE '%" + v + "%')"
		}
		str = str + ")"
	} else {
		str = "Usuario != ?"
	}
	err = orm.Where(str, session.Usuario).Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		respuesta, _ := json.Marshal(responses.Error{Success: false, Error: err.Error()})
		w.Write(respuesta)
		return
	}
	length := len(users)
	list := make([]responses.Friend, length)
	for i := 0; i < length; i++ {
		list[i] = responses.Friend{
			Usuario:   users[i].Usuario,
			Nombres:   users[i].Nombres,
			Apellidos: users[i].Apellidos,
		}
	}
	respuesta, _ := json.Marshal(responses.ListFriends{
		Success: true,
		Data:    list,
	})
	w.Write(respuesta)
}
