package groups

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/social"
	"github.com/juliotorresmoreno/unravel-server/ws"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func update(SocialBD *mgo.Database, session *models.User, ID bson.ObjectId, nombre, descripcion, permiso string, categoria int) error {
	var row = social.Group{}
	query := bson.M{
		"_id":    bson.M{"$ne": ID},
		"nombre": nombre,
	}
	num, _ := SocialBD.C(grupos).Find(query).Limit(1).Count()
	if num > 0 {
		return errors.New("El grupo ya existe")
	}
	query = bson.M{"_id": ID, "usuario": session.Usuario}
	SocialBD.C(grupos).Find(query).One(&row)
	if row.ID.String() != ID.String() {
		return errors.New("El grupo no existe")
	}
	row.Nombre = nombre
	row.Descripcion = descripcion
	row.Permiso = permiso
	var _, err = govalidator.ValidateStruct(row)
	if err != nil {
		return err
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
	return err
}

func insert(SocialBD *mgo.Database, session *models.User, nombre, descripcion, permiso string, categoria int) error {
	query := bson.M{"nombre": nombre}
	num, _ := SocialBD.C(grupos).Find(query).Limit(1).Count()
	if num > 0 {
		return errors.New("El grupo ya existe")
	}
	err := social.Add(grupos, social.Group{
		ID:          bson.NewObjectId(),
		Usuario:     session.Usuario,
		Nombre:      nombre,
		Descripcion: descripcion,
		Categoria:   categoria,
		Permiso:     permiso,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})
	return err
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
	if _ID != "" {
		err = update(SocialBD, session, ID, nombre, descripcion, permiso, categoria)
	} else {
		err = insert(SocialBD, session, nombre, descripcion, permiso, categoria)
	}
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\": true}"))
}
