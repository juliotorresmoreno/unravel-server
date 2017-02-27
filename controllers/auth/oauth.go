package auth

import (
	"encoding/json"
	"errors"
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
	autenticateOauth(w, r, content.Usuario, state)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resultado, _ := json.Marshal(content)
	w.Write(resultado)
}

func autenticateOauth(w http.ResponseWriter, r *http.Request, usuario, tipo string) {
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ? and Tipo = ?", usuario, tipo).Find(&users)
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

	helper.DespacharError(w, errors.New("Usuario o contraseña invalido"), http.StatusUnauthorized)
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
