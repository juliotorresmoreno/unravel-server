package oauth

import "net/http"

/*func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/oauth2callback", handleOauth2Callback)
	mux.HandleFunc("/auth/facebook", handleFacebook)
	mux.HandleFunc("/auth/github", handleGithub)
	mux.HandleFunc("/auth/google", handleGoogle)
	http.ListenAndServe(":80", mux)
}*/

// HandleOauth2Callback este es el calback que recive la respuesta de la autenticacion
func HandleOauth2Callback(w http.ResponseWriter, r *http.Request) {
	var state = r.FormValue("state")
	var code = r.FormValue("code")
	var content []byte
	var err error
	switch state {
	case "google":
		content, err = googleCallback(code, state)
	case "facebook":
		content, err = facebookCallback(code, state)
	case "github":
		content, err = githubCallback(code, state)
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(content)
}
