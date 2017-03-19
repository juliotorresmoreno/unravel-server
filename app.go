package main

import (
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"./config"
	"./router"
)

func main() {
	go startHTTP()
	startHTTPS()
}

func startHTTP() {
	var mux = mux.NewRouter().StrictSlash(false)
	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var url = "https://" + config.HOSTNAME + r.URL.Path + "?" + r.URL.RawQuery
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}).Methods("GET")
	var addr = ":" + strconv.Itoa(config.PORT)
	var server = &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    config.READ_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening on " + addr)
	println(server.ListenAndServe())
}

func startHTTPS() {
	var addrSsl = ":" + strconv.Itoa(config.PORT_SSL)
	var certFile = config.CERT_FILE
	var keyFile = config.KEY_FILE
	var serverSsl = &http.Server{
		Addr:           addrSsl,
		Handler:        router.GetHandler(),
		ReadTimeout:    config.READ_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening on " + addrSsl)
	println(serverSsl.ListenAndServeTLS(certFile, keyFile))
}
