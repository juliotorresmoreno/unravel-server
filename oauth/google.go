package oauth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConf = &oauth2.Config{
	ClientID:     "253896511031-dl60esqnmcedd5p0v7rgueuq5ca1nrsv.apps.googleusercontent.com",
	ClientSecret: "CgXXY8FGLRVc1wTvD89uX-1u",
	Scopes:       []string{"profile", "email", "openid"},
	RedirectURL:  "http://unravel.ga/oauth2callback",
	Endpoint:     google.Endpoint,
}

//GoogleCallback devuelve los datos basicos del usuario
func GoogleCallback(code, state string) (Usuario, error) {
	var query string
	var token, err = googleConf.Exchange(oauth2.NoContext, code)
	query = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken
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
	var resultado = map[string]interface{}{}
	var email string
	json.Unmarshal(contents, &resultado)
	if resultado["email"] == nil {
		email = fmt.Sprintf("%v", resultado["id"]) + "@google.com"
	} else {
		email = fmt.Sprintf("%v", resultado["email"])
	}
	var respuesta = Usuario{
		FullName: fmt.Sprintf("%v", resultado["name"]),
		Email:    email,
		Usuario:  fmt.Sprintf("%v", resultado["id"]) + "_google",
		Tipo:     "google",
		Code:     code,
	}
	return respuesta, nil
}

//HandleGoogle ruta de acceso a la autenticacion
func HandleGoogle(w http.ResponseWriter, r *http.Request) {
	var url = googleConf.AuthCodeURL("google", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
