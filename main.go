package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/juliotorresmoreno/unravel-server/app"
)

func main() {
	app := app.NewApp()
	app.Start()
}
