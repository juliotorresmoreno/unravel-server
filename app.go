package main

import (
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"./config"
	"./router"
)

func main() {
	var addr = ":" + strconv.Itoa(config.PORT)
	var server = &http.Server{
		Addr:           addr,
		Handler:        router.GetHandler(),
		ReadTimeout:    config.READ_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening on " + addr)
	println(server.ListenAndServe())
}
