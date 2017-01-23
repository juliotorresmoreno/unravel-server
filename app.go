package main

import (
	"net/http"
	"time"

	"./router"
)

type name struct {
	sSs string
}

func main() {
	//var methods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	//var origin = handlers.AllowedOrigins([]string{"*"})
	//var handler = handlers.CORS(methods, origin)(router.GetHandler())
	var server = &http.Server{
		Addr:           ":8080",
		Handler:        router.GetHandler(),
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening")
	println(server.ListenAndServe())
}
