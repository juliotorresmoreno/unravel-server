package groups

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

const grupos = "groups"

//ObtenerGrupos obtiene el listado de grupos
func ObtenerGrupos(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var resultado = make([]group, 0)
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query = bson.M{"usuario": session.Usuario}
	err = SocialBD.C(grupos).Find(query).All(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)
}

//ObtenerTodosGrupos obtiene el listado de grupos
func ObtenerTodosGrupos(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var resultado = make([]group, 0)
	var nombre = r.URL.Query().Get("nombre")
	if nombre != "" {
		var socialSS, SocialBD, err = social.GetSocial()
		if err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		defer socialSS.Close()
		var query = bson.M{"nombre": bson.M{"$regex": nombre}}
		err = SocialBD.C(grupos).Find(query).All(&resultado)
		if err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)
}

//Describe obtiene el listado de grupos
func Describe(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var grupo = vars["group"]
	var resultado = group{}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query = bson.M{"nombre": grupo}
	err = SocialBD.C(grupos).Find(query).One(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)
}

type group struct {
	ID          interface{} "_id"
	Usuario     string      `json:"usuario"`
	Nombre      string      `json:"nombre"`
	Descripcion string      `json:"descripcion"`
	Categoria   int         `json:"categoria"`
	Permiso     string      `json:"permiso"`
	CreateAt    time.Time   `json:"create_at"`
	UpdateAt    time.Time   `json:"update_at"`
}
