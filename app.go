package main

import (
	"net/http"
	"time"

	"./router"
	"github.com/gorilla/handlers"
)

func main() {
	var server = &http.Server{
		Addr:           ":8080",
		Handler:        handlers.CORS()(router.GetHandler()),
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening")
	println(server.ListenAndServe())
}
