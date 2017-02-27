package auth

import (
	"encoding/json"
	"net/http"

	"../../config"
	"../../helper"
	"../../models"
	"../../oauth"
)

//Oauth2Callback este es el calback que recive la respuesta de la autenticacion
func Oauth2Callback(w http.ResponseWriter, r *http.Request) {
	var state = r.FormValue("state")
	var code = r.FormValue("code")
	var content, err = obtenerDatos(code, state)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	autenticateOauth(w, r, content, state)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func autenticateOauth(w http.ResponseWriter, r *http.Request, usuario oauth.Usuario, tipo string) {
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ? and Tipo = ?", usuario.Usuario, tipo).Find(&users)
	w.Header().Set("Content-Type", "application/json")
	helper.Cors(w, r)

	if err == nil && len(users) > 0 {
		_token, _session := autenticate(&users[0])
		http.SetCookie(w, &http.Cookie{
			MaxAge:   config.SESSION_DURATION,
			Secure:   false,
			HttpOnly: true,
			Name:     "token",
			Value:    _token,
			Path:     "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ = json.Marshal(_session)
		w.Write(respuesta)
		return
	}

	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}

	registrarOauth(w, r, usuario, tipo)
}

func registrarOauth(w http.ResponseWriter, r *http.Request, usuario oauth.Usuario, tipo string) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	user.Nombres = usuario.Nombres
	user.Apellidos = usuario.Apellidos
	user.Usuario = usuario.Usuario
	user.Email = usuario.Email
	user.Code = usuario.Code
	user.Tipo = tipo
	user.Passwd = ""

	if _, err := user.Add(); err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
	} else {
		var _token, _session = autenticate(&user)
		var respuesta, _ = json.Marshal(_session)
		http.SetCookie(w, &http.Cookie{
			MaxAge:   config.SESSION_DURATION,
			HttpOnly: true,
			Secure:   false,
			Name:     "token",
			Value:    _token,
			Path:     "/",
		})
		w.WriteHeader(http.StatusCreated)
		w.Write(respuesta)
	}
}

func obtenerDatos(code, state string) (oauth.Usuario, error) {
	var content oauth.Usuario
	var err error
	switch state {
	case "google":
		content, err = oauth.GoogleCallback(code, state)
	case "facebook":
		content, err = oauth.FacebookCallback(code, state)
	case "github":
		content, err = oauth.GithubCallback(code, state)
	}
	return content, err
}
