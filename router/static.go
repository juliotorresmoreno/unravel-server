package router

import (
	"net/http"
	"os"
)

func publicHandler(w http.ResponseWriter, r *http.Request) {
	var publicPath = "./webroot"
	var path = publicPath + "/" + r.URL.Path
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Cache-Control")
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, publicPath+"/index.html")
}
