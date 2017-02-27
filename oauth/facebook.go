package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var facebookConf = &oauth2.Config{
	ClientID:     "1863527043887364",
	ClientSecret: "342afa0b1aaad4d04baf3b2867e7a992",
	Scopes:       []string{"public_profile", "user_website", "email"},
	RedirectURL:  "http://unravel.ga/oauth2callback",
	Endpoint:     facebook.Endpoint,
}

//FacebookCallback devuelve los datos basicos del usuario
func FacebookCallback(code, state string) (Usuario, error) {
	var token, err = facebookConf.Exchange(oauth2.NoContext, code)
	var query = "https://graph.facebook.com/me?access_token=" + token.AccessToken
	if err != nil {
		return Usuario{}, err
	}
	response, err := http.Get(query)
	if err != nil {
		return Usuario{}, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Usuario{}, err
	}
	var resultado map[string]interface{}
	var respuesta = Usuario{}
	json.Unmarshal(contents, resultado)
	respuesta.FullName = resultado["name"].(string)
	respuesta.Email = resultado["email"].(string)
	return respuesta, nil
}

//HandleFacebook ruta de acceso a la autenticacion
func HandleFacebook(w http.ResponseWriter, r *http.Request) {
	var url = facebookConf.AuthCodeURL("facebook", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
