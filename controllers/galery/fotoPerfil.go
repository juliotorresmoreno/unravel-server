package galery

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// GetFotoPerfil establece la foto de perfil.
func GetFotoPerfil(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var path = config.PATH + "/" + usuario + "/fotoPerfil"
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.Redirect(w, r, "/static/svg/user-3.svg", http.StatusFound)
}

// SetFotoPerfil establece la foto de perfil.
func SetFotoPerfil(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var path = config.PATH + "/" + session.Usuario
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	var name = "fotoPerfil"
	save, _ := os.Create(path + "/" + name)
	defer save.Close()
	_, err = io.Copy(save, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\": true}"))
}
