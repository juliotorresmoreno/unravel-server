package galery

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
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
	http.Redirect(w, r, "/icons/148705-essential-collection/png/picture-2.png", http.StatusFound)
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

	var ramd = helper.GenerateRandomString(20)
	var tnme = "/tmp/" + ramd + ".tmp"
	tmp, err := os.Create(tnme)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer func() {
		tmp.Close()
		os.Remove(tnme)
	}()

	if _, err = io.Copy(tmp, file); err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}

	name := path + "/" + "fotoPerfil"
	helper.BuildJPG(tnme, name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\": true}"))
}
