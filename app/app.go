package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/router"
)

type App struct {
	http.Handler
}

func NewApp() App {
	app := App{
		Handler: router.NewRouter(),
	}
	return app
}

func (app App) Start() {
	go app.ListenHttp()
	<-make(chan int)
}

func (app App) ListenHttp() {
	addr := fmt.Sprintf(":%v", config.PORT)
	server := http.Server{
		Addr:        addr,
		Handler:     app.Handler,
		ReadTimeout: 10 * time.Second,
	}
	fmt.Printf("Listen on %v\n", addr)
	server.ListenAndServe()
}
