package galery

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// Upload sube las imagenes
func Upload(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var galery = strings.Trim(r.FormValue("galery"), " ")
	var galeria = config.PATH + "/" + session.Usuario + "/" + galery + "/images"
	var mini = config.PATH + "/" + session.Usuario + "/" + galery + "/mini"
	os.MkdirAll(galeria, 0755)
	os.MkdirAll(mini, 0755)

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	var ramd = helper.GenerateRandomString(20)
	var name = galeria + "/" + ramd + ".jpg"
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
	_, err = io.Copy(tmp, file)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	helper.BuildJPG(tnme, name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\": true}"))
}
