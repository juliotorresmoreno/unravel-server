package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"../controllers/auth"
	"../test"
)

func GetHandler() http.Handler {
	webroot := http.FileServer(http.Dir("webroot"))
	mux := mux.NewRouter().StrictSlash(false)
	mux.Handle("/", webroot)
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")

	mux.HandleFunc("/test", test.Test).Methods("GET")
	return mux
}

//func (e mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, e.msg)
//}

/*mux.Handle("/api/users", getUsers).Methods("GET")
mux.Handle("/api/users", postUsers).Methods("POST")
mux.Handle("/api/users", putUsers).Methods("PUT")
mux.Handle("/api/users", deleteUsers).Methods("DELETE")*/