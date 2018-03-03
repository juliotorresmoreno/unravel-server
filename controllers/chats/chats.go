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
