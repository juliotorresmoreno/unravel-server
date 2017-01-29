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

const noticias = "noticias"

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var noticia = r.PostFormValue("noticia")
	var permiso = r.PostFormValue("permiso")
	if !helper.IsValidPermision(permiso) {
		despacharError(w, errors.New("Permiso denegado"), http.StatusBadRequest)
		return
	}
	var nueva = &social.Noticia{
		Usuario:  session.Usuario,
		Noticia:  noticia,
		Likes:    make([]string, 0),
		Permiso:  permiso,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	var err = social.Add(noticias, nueva)
	if err != nil {
		despacharError(w, err, http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\":true}"))
}

func despacharError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	})
	w.Write(respuesta)
}

// Listar listado de noticias en el muro
func Listar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var query = bson.M{"usuario": session.Usuario}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var resultado = make([]noticia, 0)
	err = SocialBD.C(noticias).Find(query).Sort("-createat").All(&resultado)
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
		return
	}
	var _usuarios = make([]string, 0)
	for i := 0; i < len(resultado); i++ {
		_usuarios = append(_usuarios, resultado[i].Usuario)
		for j := 0; j < len(resultado[i].Comentarios); j++ {
			_usuarios = append(_usuarios, resultado[i].Comentarios[j].Usuario)
		}
	}
	usuarios, err := models.FindUsers(_usuarios)
	for _, value := range usuarios {
		for i := 0; i < len(resultado); i++ {
			if resultado[i].Usuario == value.Usuario {
				resultado[i].Nombres = value.Nombres
				resultado[i].Apellidos = value.Apellidos
			}
			for j := 0; j < len(resultado[i].Comentarios); j++ {
				if resultado[i].Comentarios[j].Usuario == value.Usuario {
					resultado[i].Comentarios[j].Nombres = value.Nombres
					resultado[i].Comentarios[j].Apellidos = value.Apellidos
				}
			}
		}
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
	var ID = bson.ObjectIdHex(r.PostFormValue("noticia"))
	var query = map[string]interface{}{
		"_id":     ID,
		"usuario": session.Usuario,
	}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var resultado = noticia{}
	err = SocialBD.C(noticias).Find(query).One(&resultado)
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
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
		despacharError(w, err, http.StatusInternalServerError)
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
	if r.PostFormValue("noticia") == "" {
		despacharError(w, errors.New("Falta la noticia"), http.StatusNotAcceptable)
		return
	}
	if r.PostFormValue("comentario") == "" {
		despacharError(w, errors.New("Falta el comentario"), http.StatusNotAcceptable)
		return
	}
	var ID = bson.ObjectIdHex(r.PostFormValue("noticia"))
	var _comentario = r.PostFormValue("comentario")
	var query = bson.M{
		"_id":     ID,
		"usuario": session.Usuario,
	}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var resultado = noticia{}
	err = SocialBD.C(noticias).Find(query).One(&resultado)
	if err != nil {
		despacharError(w, err, http.StatusInternalServerError)
		return
	}
	resultado.Comentarios = append(resultado.Comentarios, comentario{
		Comentario: _comentario,
		Usuario:    session.Usuario,
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
		despacharError(w, err, http.StatusInternalServerError)
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
