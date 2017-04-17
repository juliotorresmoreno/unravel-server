package galery

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// ViewImagen ver imagen
func ViewImagen(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var imagen = vars["imagen"]
	var mini = r.URL.Query().Get("mini")
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	if mini != "" {
		var path = config.PATH + "/" + usuario + "/" + galeria + "/mini/" + imagen
		var source = config.PATH + "/" + usuario + "/" + galeria + "/images/" + imagen
		if f, err := os.Stat(path); err == nil && !f.IsDir() {
			http.ServeFile(w, r, path)
			return
		}
		helper.BuildMini(source, path)
		http.ServeFile(w, r, path)
		return
	}
	var path = config.PATH + "/" + usuario + "/" + galeria + "/images/" + imagen
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}

// ViewPreview ver preview de galeria
func ViewPreview(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var usuario, imagen, auth, url string
	var defecto = "/static/svg/user-3.svg"
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var token = helper.GetToken(r)
	var imagenes = listarImagenes(usuario, galeria)
	var length = len(imagenes)
	if length == 0 {
		url = "https://" + r.Host + defecto
	} else {
		imagen = imagenes[rand.Intn(length)]
		auth = "?token=" + token
		url = "https://" + r.Host + "/api/v1/" + usuario + "/galery/" + galeria + "/" + imagen + auth + "&mini=1"
	}
	http.Redirect(w, r, url, http.StatusFound)
}
