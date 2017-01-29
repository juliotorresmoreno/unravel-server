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
		despacharError(w, errors.New("Permiso denegado"))
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
		despacharError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\":true}"))
}

func despacharError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
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
		despacharError(w, err)
		return
	}
	defer socialSS.Close()
	var resultado = make([]noticia, 0)
	err = SocialBD.C(noticias).Find(query).Sort("-createat").All(&resultado)
	if err != nil {
		despacharError(w, err)
		return
	}
	var length = len(resultado)
	var _usuarios = make([]string, length)
	for i := 0; i < length; i++ {
		_usuarios[i] = resultado[i].Usuario
	}
	usuarios, err := models.FindUsers(_usuarios)
	for i := 0; i < length; i++ {
		for _, value := range usuarios {
			if resultado[i].Usuario == value.Usuario {
				resultado[i].Nombres = value.Nombres
				resultado[i].Apellidos = value.Apellidos
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
	var ID = r.PostFormValue("noticia")
	var query = bson.M{"usuario": session.Usuario}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		despacharError(w, err)
		return
	}
	defer socialSS.Close()
	var resultado = noticia{}
	err = SocialBD.C(noticias).Find(query).One(&resultado)
	if err != nil {
		despacharError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var existe = find(resultado.Likes, session.Usuario)
	if existe >= 0 {
		resultado.Likes = remove(resultado.Likes, existe)
	} else {
		resultado.Likes = append(resultado.Likes, session.Usuario)
	}
	err = SocialBD.C(noticias).UpdateId(ID, resultado)
	if err != nil {
		despacharError(w, err)
		return
	}
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
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

type noticia struct {
	ID        interface{} "_id"
	Usuario   string      `json:"usuario"`
	Nombres   string      `json:"nombres"`
	Apellidos string      `json:"apellidos"`
	Noticia   string      `json:"noticia"`
	Permiso   string      `json:"permiso"`
	Likes     []string    `json:"likes"`
	CreateAt  time.Time   `json:"create_at"`
	UpdateAt  time.Time   `json:"update_at"`
}
