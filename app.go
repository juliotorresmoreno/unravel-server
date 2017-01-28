package main

import "net/http"
import "time"
import _ "github.com/go-sql-driver/mysql"
import "./router"
import "./config"

func main() {
	var server = &http.Server{
		Addr:           ":" + string(config.PORT),
		Handler:        router.GetHandler(),
		ReadTimeout:    config.READ_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening")
	println(server.ListenAndServe())
}
