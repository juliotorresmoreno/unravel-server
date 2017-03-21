package news

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"../../helper"
	"../../models"
	"../../social"
	"../../ws"
	"gopkg.in/mgo.v2/bson"
)

// Comentar el de toda la vida
func Comentar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	defer func() {
		if r := recover(); r != nil {
			helper.DespacharError(w, errors.New("Error desconocido"), http.StatusInternalServerError)
		}
	}()
	if r.PostFormValue("noticia") == "" {
		helper.DespacharError(w, errors.New("Falta la noticia"), http.StatusNotAcceptable)
		return
	}
	if r.PostFormValue("comentario") == "" {
		helper.DespacharError(w, errors.New("Falta el comentario"), http.StatusNotAcceptable)
		return
	}
	var ID = bson.ObjectIdHex(r.PostFormValue("noticia"))
	var _comentario = r.PostFormValue("comentario")
	var friends, _ = models.GetFriends(session.Usuario)
	var length = len(friends)
	var usuarios = make([]string, length+1)
	for i := 0; i < length; i++ {
		usuarios[i] = friends[i].Usuario
	}
	usuarios[length] = session.Usuario
	var query = bson.M{
		"_id": ID,
		"usuario": map[string]interface{}{
			"$in": usuarios,
		},
	}

	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var resultado = noticia{}
	err = SocialBD.C(noticias).Find(query).One(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	resultado.Comentarios = append(resultado.Comentarios, comentario{
		Usuario:    session.Usuario,
		Nombres:    session.Nombres,
		Apellidos:  session.Apellidos,
		FullName:   session.FullName,
		Comentario: _comentario,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	})
	var data = map[string]interface{}{
		"$set": map[string]interface{}{
			"comentarios": resultado.Comentarios,
			"updateat":    time.Now(),
		},
	}
	err = SocialBD.C(noticias).UpdateId(ID, data)
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
