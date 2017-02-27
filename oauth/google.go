package oauth

import (
	"io/ioutil"
	"net/http"

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
func GoogleCallback(code, state string) ([]byte, error) {
	var query string
	var token, err = googleConf.Exchange(oauth2.NoContext, code)
	query = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken
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

//HandleGoogle ruta de acceso a la autenticacion
func HandleGoogle(w http.ResponseWriter, r *http.Request) {
	var url = googleConf.AuthCodeURL("google", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
