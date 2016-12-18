package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"../controllers/auth"
	"../controllers/friends"
	"../test"
)

func GetHandler() http.Handler {
	webroot := http.FileServer(http.Dir("webroot"))
	mux := mux.NewRouter().StrictSlash(false)

	// auth
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")
	mux.HandleFunc("/api/v1/auth/login", auth.Login).Methods("POST")
	mux.HandleFunc("/api/v1/auth/session", auth.Session).Methods("GET")

	mux.HandleFunc("/api/v1/friends", friends.ListFriends).Methods("GET")

	mux.HandleFunc("/test", test.Test).Methods("GET")

	mux.Handle("/", webroot)
	mux.Handle("/{all}", webroot)
	return mux
}
