package news

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"../../helper"
	"../../models"
	"../../social"
	"../../ws"
	"gopkg.in/mgo.v2/bson"
)

const noticias = "noticias"

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

// Like el like de toda la vida
func Like(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	defer func() {
		if r := recover(); r != nil {
			helper.DespacharError(w, errors.New("Error desconocido"), http.StatusInternalServerError)
		}
	}()
	var ID = bson.ObjectIdHex(r.PostFormValue("noticia"))
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
	var existe = find(resultado.Likes, session.Usuario)
	if existe >= 0 {
		resultado.Likes = remove(resultado.Likes, existe)
	} else {
		resultado.Likes = append(resultado.Likes, session.Usuario)
	}
	var data = map[string]interface{}{
		"$set": map[string]interface{}{
			"likes":    resultado.Likes,
			"updateat": time.Now(),
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
	Comentario string    `json:"comentarios"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
