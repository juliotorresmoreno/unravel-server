package news

import (
	"errors"
	"net/http"
	"time"

	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
	"github.com/unravel-server/social"
	"github.com/unravel-server/ws"
	"gopkg.in/mgo.v2/bson"
)

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var _noticia = r.PostFormValue("noticia")
	var _permiso = r.PostFormValue("permiso")
	if !helper.IsValidPermision(_permiso) {
		helper.DespacharError(w, errors.New("Permiso denegado"), http.StatusBadRequest)
		return
	}
	var nueva = &social.Noticia{
		ID:        bson.NewObjectId(),
		Usuario:   session.Usuario,
		Nombres:   session.Nombres,
		Apellidos: session.Apellidos,
		FullName:  session.FullName,
		Noticia:   _noticia,
		Likes:     make([]string, 0),
		Permiso:   _permiso,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	var err = social.Add(noticias, nueva)
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\":true}"))
}
