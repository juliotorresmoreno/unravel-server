package router

import (
	"net/http"
	"os"
)

func publicHandler(w http.ResponseWriter, r *http.Request) {
	var publicPath = "./webroot"
	var path = publicPath + "/" + r.URL.Path
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, publicPath+"/index.html")
}
