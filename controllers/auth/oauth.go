package auth

import "net/http"
import "../../oauth"
import "encoding/json"

//Oauth2Callback este es el calback que recive la respuesta de la autenticacion
func Oauth2Callback(w http.ResponseWriter, r *http.Request) {
	var state = r.FormValue("state")
	var code = r.FormValue("code")
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
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resultado, _ := json.Marshal(content)
	w.Write(resultado)
}
