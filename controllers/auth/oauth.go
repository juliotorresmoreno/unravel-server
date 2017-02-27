package auth

import (
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
	err = autenticateOauth(w, content, state)
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}
}

func autenticateOauth(w http.ResponseWriter, usuario oauth.Usuario, tipo string) error {
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ? and Tipo = ?", usuario.Usuario, tipo).Find(&users)

	if err == nil && len(users) > 0 {
		_token, _ := autenticate(&users[0])
		var cookie = http.Cookie{
			MaxAge:   config.SESSION_DURATION,
			HttpOnly: true,
			Secure:   false,
			Name:     "token",
			Value:    _token,
			Path:     "/",
		}
		w.Header().Set("Location", "/")
		w.Header().Set("Set-Cookie", cookie.String())
		w.WriteHeader(http.StatusTemporaryRedirect)
		return nil
	}

	if err != nil {
		return err
	}
	return registrarOauth(w, usuario, tipo)
}

func registrarOauth(w http.ResponseWriter, usuario oauth.Usuario, tipo string) error {
	var user models.User
	user.Nombres = usuario.Nombres
	user.Apellidos = usuario.Apellidos
	user.FullName = usuario.FullName
	user.Usuario = usuario.Usuario
	user.Email = usuario.Email
	user.Code = usuario.Code
	user.Tipo = tipo
	user.Passwd = ""

	if _, err := user.ForceAdd(); err != nil {
		return err
	}
	var _token, _ = autenticate(&user)
	var cookie = http.Cookie{
		MaxAge:   config.SESSION_DURATION,
		HttpOnly: true,
		Secure:   false,
		Name:     "token",
		Value:    _token,
		Path:     "/",
	}
	w.Header().Set("Location", "/")
	w.Header().Set("Set-Cookie", cookie.String())
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
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
