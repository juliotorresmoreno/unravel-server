package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubConf = &oauth2.Config{
	ClientID:     "c572b3942abce89a0c57",
	ClientSecret: "dced921d29e2cda80890b13c8bf833527f3e2b9d",
	Scopes:       []string{"user", "public_repo"},
	RedirectURL:  "http://unravel.ga/oauth2callback",
	Endpoint:     github.Endpoint,
}

//GithubCallback devuelve los datos basicos del usuario
func GithubCallback(code, state string) (Usuario, error) {
	var query string
	var token, err = githubConf.Exchange(oauth2.NoContext, code)
	query = "https://api.github.com/user?access_token=" + token.AccessToken
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
	json.Unmarshal(contents, &resultado)
	var respuesta = Usuario{
		FullName: fmt.Sprintf("%v", resultado["name"]),
		Email:    fmt.Sprintf("%v@github.com", resultado["id"]),
		Usuario:  fmt.Sprintf("%v_github", resultado["login"]),
		Tipo:     "github",
		Code:     code,
	}
	return respuesta, nil
}

//HandleGithub ruta de acceso a la autenticacion
func HandleGithub(w http.ResponseWriter, r *http.Request) {
	var url = githubConf.AuthCodeURL("github", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
