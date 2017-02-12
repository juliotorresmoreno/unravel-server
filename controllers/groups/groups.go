package groups

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	"strconv"

	"../../config"
	"../../helper"
	"../../models"
	"../../social"
	"../../ws"
	"gopkg.in/mgo.v2/bson"
)

const grupos = "groups"

//Preview Previsualizacion de la miniatura del grupo
func Preview(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var grupo = vars["group"]
	var row = social.Group{}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query = bson.M{"nombre": grupo}
	SocialBD.C(grupos).Find(query).One(&row)
	if row.Usuario == "" {
		helper.DespacharError(w, errors.New("El grupo no existe"), http.StatusNotAcceptable)
		return
	}
	var usuario = row.Usuario
	var galery = "groups"
	var galeria = config.PATH + "/" + usuario + "/" + galery + "/images"
	var path = galeria + "/" + row.ID.Hex() + ".jpg"
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.Redirect(w, r, "http://semantic-ui.com/images/wireframe/image.png", http.StatusFound)
}

//ChangePreview cambia el preview del grupo
func ChangePreview(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var group = strings.Trim(r.FormValue("group"), " ")
	var galery = "groups"
	var galeria = config.PATH + "/" + session.Usuario + "/" + galery + "/images"
	os.MkdirAll(galeria, 0755)

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	var ramd = helper.GenerateRandomString(20)
	var name = galeria + "/" + group + ".jpg"
	var tnme = "/tmp/" + ramd + ".tmp"

	tmp, err := os.Create(tnme)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer func() {
		tmp.Close()
		os.Remove(tnme)
	}()
	_, err = io.Copy(tmp, file)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	helper.BuildMini(tnme, name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}

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
	var categoria, _ = strconv.Atoi(r.PostFormValue("categoria"))
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
		query = bson.M{
			"_id":    bson.M{"$ne": ID},
			"nombre": nombre,
		}
		num, _ := SocialBD.C(grupos).Find(query).Limit(1).Count()
		if num > 0 {
			helper.DespacharError(w, errors.New("El grupo ya existe"), http.StatusNotAcceptable)
			return
		}
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
				"categoria":   categoria,
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
			Categoria:   categoria,
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
	Categoria   int         `json:"categoria"`
	Permiso     string      `json:"permiso"`
	CreateAt    time.Time   `json:"create_at"`
	UpdateAt    time.Time   `json:"update_at"`
}
