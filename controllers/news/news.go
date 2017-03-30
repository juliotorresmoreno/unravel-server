package news

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
	"github.com/unravel-server/social"
	"github.com/unravel-server/ws"
	"gopkg.in/mgo.v2/bson"
)

const noticias = "noticias"

// GetNews listado de noticias en el muro
func GetNews(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var antesDe = r.URL.Query().Get("antesDe")
	var despuesDe = r.URL.Query().Get("despuesDe")
	var query = bson.M{"usuario": usuario}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	if vars["usuario"] == "" {
		var friends, _ = models.GetFriends(session.Usuario)
		var _friends = make([]string, len(friends)+1)
		_friends[len(friends)] = session.Usuario
		for i, v := range friends {
			_friends[i] = v.Usuario
		}
		query["usuario"] = map[string]interface{}{
			"$in": _friends,
		}
	}
	if antesDe != "" {
		tiempo, _ := time.Parse(time.RFC3339, antesDe)
		query["createat"] = bson.M{
			"$lt": tiempo,
		}
	}
	if despuesDe != "" {
		tiempo, _ := time.Parse(time.RFC3339, despuesDe)
		query["createat"] = bson.M{
			"$gt": tiempo,
		}
	}
	var resultado = make([]noticia, 0)
	err = SocialBD.C(noticias).Find(query).Sort("-createat").Limit(10).All(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Write(respuesta)
}

func find(arr []string, cadena string) int {
	for i, v := range arr {
		if v == cadena {
			return i
		}
	}
	return -1
}

func remove(s []string, i int) []string {
	var length = len(s) - 1
	s[length], s[i] = s[i], s[length]
	return s[:length]
}

type noticia struct {
	ID          interface{}  "_id"
	Usuario     string       `json:"usuario"`
	Nombres     string       `json:"nombres"`
	Apellidos   string       `json:"apellidos"`
	FullName    string       `json:"fullname"`
	Noticia     string       `json:"noticia"`
	Permiso     string       `json:"permiso"`
	Likes       []string     `json:"likes"`
	Comentarios []comentario `json:"comentarios"`
	CreateAt    time.Time    `json:"create_at"`
	UpdateAt    time.Time    `json:"update_at"`
}

type comentario struct {
	Usuario    string    `json:"usuario"`
	Nombres    string    `json:"nombres"`
	Apellidos  string    `json:"apellidos"`
	FullName   string    `json:"fullname"`
	Comentario string    `json:"comentarios"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
