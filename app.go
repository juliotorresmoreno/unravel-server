package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type mensaje struct {
	msg string
}

func (e mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, e.msg)
}

func main() {
	getUsers := mensaje{msg: "get"}
	postUsers := mensaje{msg: "post"}
	putUsers := mensaje{msg: "put"}
	deleteUsers := mensaje{msg: "delete"}
	mux := mux.NewRouter().StrictSlash(false)
	webroot := http.FileServer(http.Dir("Webroot"))
	mux.Handle("/", webroot)
	mux.Handle("/api/users", getUsers).Methods("GET")
	mux.Handle("/api/users", postUsers).Methods("POST")
	mux.Handle("/api/users", putUsers).Methods("PUT")
	mux.Handle("/api/users", deleteUsers).Methods("DELETE")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening")
	log.Println(server.ListenAndServe())
}

func serverHTTP() {
	fs := http.FileServer(http.Dir("webroot"))
	home := mensaje{msg: "home"}
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.Handle("/home", home)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening")
	log.Println(server.ListenAndServe())
}
