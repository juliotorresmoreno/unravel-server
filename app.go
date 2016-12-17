package main

import (
	"log"
	"net/http"
	"time"

	"./router"
)

func main() {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router.GetHandler(),
		ReadTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening")
	log.Println(server.ListenAndServe())
}
