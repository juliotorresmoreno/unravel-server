package chats

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

type chat struct {
	UsuarioEmisor   string `json:"usuario_emisor"`
	UsuarioReceptor string `json:"usuario_receptor"`
}

// TableName establece el nombre de la tabla del modelo
func (el chat) TableName() string {
	return "chats"
}

// User modelo de usuario
type user struct {
	ID        uint   `json:"id" xorm:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	FullName  string `json:"fullname"`
	Usuario   string `json:"usuario"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (el user) TableName() string {
	return "users"
}

//GetAll d
func GetAll(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	var orm = db.GetXORM()
	defer orm.Close()

	usuario := session.Usuario
	chats := make([]chat, 0)
	err := orm.
		Distinct("usuario_emisor,usuario_receptor").
		Where("usuario_emisor = ? OR usuario_receptor = ?", usuario, usuario).
		Find(&chats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	length := len(chats)
	users := make([]user, 0, length)
	exists := map[string]bool{}
	for i := 0; i < length; i++ {
		usuarioChat := chats[i].UsuarioEmisor
		if usuarioChat == usuario {
			usuarioChat = chats[i].UsuarioReceptor
		}
		if _, ok := exists[usuarioChat]; ok {
			continue
		}
		exists[usuarioChat] = true
		user := user{}
		if _, err := orm.Where("usuario = ?", usuarioChat).Get(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

// GetConversacion obtiene la conversacion con el usuario solicitado
func GetConversacion(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	var orm = db.GetXORM()
	defer orm.Close()
	var resultado = make([]models.Chat, 0)
	var usuario = vars["user"]
	var antesDe = r.URL.Query().Get("antesDe")
	var despuesDe = r.URL.Query().Get("despuesDe")

	var cond = "(usuario_receptor = ? and usuario_emisor = ?) or (usuario_receptor = ? and usuario_emisor = ?)"

	var updater *xorm.Session
	var consultor *xorm.Session
	var leido = models.Chat{Leido: 1}
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
