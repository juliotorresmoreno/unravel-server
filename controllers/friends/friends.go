package friends

import (
	"net/http"
	"encoding/json"
	"../../models"
	"../responses"
)

func ListFriends(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Find(&users)
	if err != nil {
		respuesta, _ := json.Marshal(responses.Error{Success:false,Error:err.Error()})
		w.Write(respuesta)
		return
	}
	length := len(users)
	list := make([]responses.Friend, length)
	for i := 0; i < length; i++ {
		list[i] = responses.Friend{
			Usuario: users[i].Usuario,
			Nombres: users[i].Nombres,
			Apellidos: users[i].Apellidos,
		}
	}
	respuesta, _ := json.Marshal(responses.ListFriends{
		Success: true,
		Data: list,
	})
	w.Write(respuesta)
}