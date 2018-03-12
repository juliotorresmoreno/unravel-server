package users

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

type luser struct {
	models.Friend
	Legenda     string `json:"legenda"`
	Descripcion string `json:"descripcion"`
}

// Find buscar usuarios
func Find(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	query := r.URL.Query().Get("q")
	user := r.URL.Query().Get("u")
	users, _ := models.FindUser(session.Usuario, query, user)
	if len(users) == 1 && users[0].Relacion != nil && users[0].Relacion.EstadoRelacion == models.EstadoAceptado {
		users[0].Conectado = hub.IsConnect(users[0].Usuario)
	}
	length := len(users)
	data := make([]luser, length)
	var orm = db.GetXORM()
	defer orm.Close()
	for index := 0; index < length; index++ {
		profile := models.Profile{}
		orm.Where("usuario = ?", users[index].Usuario).Get(&profile)
		data[index].Legenda = profile.Legenda
		data[index].Descripcion = profile.Descripcion
		data[index].Friend = *users[index]
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    data,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
