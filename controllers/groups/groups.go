package groups

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	"../../helper"
	"../../models"
	"../../social"
	"../../ws"
	"gopkg.in/mgo.v2/bson"
)

const grupos = "groups"

//Save guardar el grupo
func Save(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	defer func() {
		if r := recover(); r != nil {
			helper.DespacharError(w, errors.New("Error desconocido"), http.StatusInternalServerError)
		}
	}()
	var _ID = r.PostFormValue("ID")
	var ID bson.ObjectId
	if _ID != "" {
		ID = bson.ObjectIdHex(_ID)
	}
	var nombre = r.PostFormValue("nombre")
	var descripcion = r.PostFormValue("descripcion")
	var permiso = r.PostFormValue("permiso")
	if !helper.IsValidPermision(permiso) {
		helper.DespacharError(w, errors.New("Permiso invalido"), http.StatusNotAcceptable)
		return
	}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query bson.M
	if _ID != "" {
		var row = social.Group{}
		query = bson.M{"_id": ID, "usuario": session.Usuario}
		SocialBD.C(grupos).Find(query).One(&row)
		if row.ID.String() != ID.String() {
			helper.DespacharError(w, errors.New("El grupo no existe"), http.StatusNotAcceptable)
			return
		}
		row.Nombre = nombre
		row.Descripcion = descripcion
		row.Permiso = permiso
		var _, err = govalidator.ValidateStruct(row)
		if err != nil {
			helper.DespacharError(w, err, http.StatusNotAcceptable)
			return
		}
		var data = map[string]interface{}{
			"$set": map[string]interface{}{
				"nombre":      nombre,
				"descripcion": descripcion,
				"permiso":     permiso,
				"updateat":    time.Now(),
			},
		}
		err = SocialBD.C(grupos).UpdateId(ID, data)
	} else {
		query = bson.M{"nombre": nombre}
		num, _ := SocialBD.C(grupos).Find(query).Limit(1).Count()
		if num > 0 {
			helper.DespacharError(w, errors.New("El grupo ya existe"), http.StatusNotAcceptable)
			return
		}
		err = social.Add(grupos, social.Group{
			Usuario:     session.Usuario,
			Nombre:      nombre,
			Descripcion: descripcion,
			Permiso:     permiso,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		})
	}
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
	Permiso     string      `json:"permiso"`
	CreateAt    time.Time   `json:"create_at"`
	UpdateAt    time.Time   `json:"update_at"`
}
