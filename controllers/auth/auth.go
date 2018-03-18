package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/controllers/auth/oauth"
	"github.com/juliotorresmoreno/unravel-server/middlewares"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func NewRouter(hub *ws.Hub) http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/registrar", Registrar).Methods("POST")
	mux.HandleFunc("/login", Login).Methods("POST")
	mux.HandleFunc("/session", middlewares.Protect(Session, hub, false)).Methods("GET")
	mux.HandleFunc("/logout", Logout).Methods("GET")
	mux.HandleFunc("/recovery", Recovery).Methods("POST")
	mux.HandleFunc("/password", Password).Methods("POST")
	mux.HandleFunc("/password_change", middlewares.Protect(PasswordChange, hub, true)).Methods("POST")
	mux.HandleFunc("/oauth2callback", Oauth2Callback).Methods("GET")
	mux.HandleFunc("/facebook", oauth.HandleFacebook).Methods("GET")
	mux.HandleFunc("/github", oauth.HandleGithub).Methods("GET")
	mux.HandleFunc("/google", oauth.HandleGoogle).Methods("GET")

	return mux
}
