package friends

import (
	"encoding/json"
	"net/http"

	"../../models"
	"../../ws"
	"../responses"
)

// ListFriends listado de amigos o personas con las que se puede chatear
func ListFriends(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users := make([]models.User, 0)
	orm := models.GetXORM()

	err := orm.Where("Usuario != ?", session.Usuario).Find(&users)
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
