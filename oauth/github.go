package oauth

import (
	"io/ioutil"
	"net/http"

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
func GithubCallback(code, state string) ([]byte, error) {
	var query string
	var token, err = githubConf.Exchange(oauth2.NoContext, code)
	query = "https://api.github.com/user?access_token=" + token.AccessToken
	if err != nil {
		return nil, err
	}
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	return contents, err
}

//HandleGithub ruta de acceso a la autenticacion
func HandleGithub(w http.ResponseWriter, r *http.Request) {
	var url = githubConf.AuthCodeURL("github", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
