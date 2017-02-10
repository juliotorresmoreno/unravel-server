package groups

import (
	"errors"
	"net/http"
	"time"

	"encoding/json"

	"../../helper"
	"../../models"
	"../../social"
	"../../ws"
	"gopkg.in/mgo.v2/bson"
)

const grupos = "groups"

//Save guardar el grupo
func Save(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var nombre = r.PostFormValue("nombre")
	var descripcion = r.PostFormValue("descripcion")
	var permiso = r.PostFormValue("permiso")
	if !helper.IsValidPermision(permiso) {
		helper.DespacharError(w, errors.New("Permiso invalido. "), http.StatusNotAcceptable)
		return
	}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query = bson.M{
		"usuario": session.Usuario,
		"nombre":  nombre,
	}
	num, _ := SocialBD.C(grupos).Find(query).Limit(1).Count()
	if num > 0 {
		helper.DespacharError(w, errors.New("El grupo ya existe"), http.StatusNotAcceptable)
		return
	}

	err = social.Add("groups", social.Group{
		Usuario:     session.Usuario,
		Nombre:      nombre,
		Descripcion: descripcion,
		Permiso:     permiso,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\": true}"))
}

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

type group struct {
	ID          interface{} "_id"
	Usuario     string      `json:"usuario"`
	Nombre      string      `json:"nombre"`
	Descripcion string      `json:"descripcion"`
	Permiso     string      `json:"permiso"`
	CreateAt    time.Time   `json:"create_at"`
	UpdateAt    time.Time   `json:"update_at"`
}
