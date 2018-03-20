package chats

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/middlewares"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func NewRouter(hub *ws.Hub) http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/mensaje", middlewares.Protect(MensajeAdd, hub, true)).Methods("POST")
	mux.HandleFunc("/mensaje", middlewares.Protect(MensajeEdit, hub, true)).Methods("PUT")
	mux.HandleFunc("/videollamada", middlewares.Protect(VideoLlamada, hub, true)).Methods("POST")
	mux.HandleFunc("/rechazarvideollamada", middlewares.Protect(RechazarVideoLlamada, hub, true)).Methods("POST")
	mux.HandleFunc("/{user}", middlewares.Protect(GetConversacion, hub, true)).Methods("GET")
	mux.HandleFunc("/", middlewares.Protect(GetAll, hub, true)).Methods("GET")

	return mux
}

type chat struct {
	UsuarioEmisor   string `json:"usuario_emisor"`
	UsuarioReceptor string `json:"usuario_receptor"`
}

// TableName establece el nombre de la tabla del modelo
func (el chat) TableName() string {
	return "chats"
}

// User modelo de usuario
type User struct {
	ID        uint   `json:"id" xorm:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	FullName  string `json:"fullname"`
	Usuario   string `json:"usuario"`
}

type userChat struct {
	User
	Count int64 `json:"count"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (el User) TableName() string {
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
	users := make([]userChat, 0, length)
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
		user := User{}
		if _, err := orm.Where("usuario = ?", usuarioChat).Get(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		where := "usuario_receptor = ? and usuario_emisor = ? and leido = 0"
		total, err := orm.Where(where, usuario, usuarioChat).Count(chat{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		users = append(users, userChat{User: user, Count: total})
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

	updater := orm.NewSession()
	consultor := orm.NewSession()
	leido := models.Chat{Leido: 1}

	cond := "usuario_receptor = ? and usuario_emisor = ?"
	updater = updater.Where(cond, usuario, session.Usuario)
	consultor = consultor.Where(cond, usuario, session.Usuario).
		Or(cond, session.Usuario, usuario)
	videoCall := consultor.Where("message = '@chats/videocall' and status in (1, 4)")
	if antesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, antesDe)
		tiempo := tmp.String()[0:19]

		updater = updater.And("create_at < ?", tiempo)
		consultor = consultor.And("create_at < ?", tiempo)
	} else if despuesDe != "" {
		tmp, _ := time.Parse(time.RFC3339, despuesDe)
		tiempo := tmp.String()[0:19]
		updater = updater.And("create_at < ?", tiempo)
		consultor = consultor.And("create_at > ?", tiempo)
	}
	count, err := videoCall.Count(models.Chat{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	updater.Cols("leido").Update(leido)
	consultor.Limit(10).Asc("id").Find(&resultado)
	length := len(resultado)
	conversacion := make([]map[string]interface{}, length)
	if length > 0 {
		for i := 0; i < length; i++ {
			conversacion[i] = map[string]interface{}{
				"action":          "mensaje",
				"id":              resultado[i].Id,
				"status":          resultado[i].Status,
				"usuario":         resultado[i].UsuarioEmisor,
				"usuarioReceptor": resultado[i].UsuarioReceptor,
				"mensaje":         resultado[i].Message,
				"fecha":           resultado[i].CreateAt,
			}
		}
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success":   true,
		"data":      conversacion,
		"videoCall": count > int64(0),
	})
	w.Write(respuesta)
}
