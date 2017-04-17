package router

import (
	"net/http"
	"os"

	"github.com/juliotorresmoreno/unravel-server/helper"
)

func publicHandler(w http.ResponseWriter, r *http.Request) {
	var publicPath = "./webroot"
	var path = publicPath + r.URL.Path
	helper.Cors(w, r)
	if f, err := os.Stat(path); err == nil && !f.IsDir() && path != "/index.html" {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, publicPath)
}
