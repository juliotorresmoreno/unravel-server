package groups

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/unravel-server/config"
	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
	"github.com/unravel-server/social"
	"github.com/unravel-server/ws"
	"gopkg.in/mgo.v2/bson"
)

//Preview Previsualizacion de la miniatura del grupo
func Preview(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var grupo = vars["group"]
	var row = social.Group{}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var query = bson.M{"nombre": grupo}
	SocialBD.C(grupos).Find(query).One(&row)
	if row.Usuario == "" {
		helper.DespacharError(w, errors.New("El grupo no existe"), http.StatusNotAcceptable)
		return
	}
	var usuario = row.Usuario
	var galery = "groups"
	var galeria = config.PATH + "/" + usuario + "/" + galery + "/images"
	var path = galeria + "/" + row.ID.Hex() + ".jpg"
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.Redirect(w, r, "http://semantic-ui.com/images/wireframe/image.png", http.StatusFound)
}

//ChangePreview cambia el preview del grupo
func ChangePreview(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var group = strings.Trim(r.FormValue("group"), " ")
	var galery = "groups"
	var galeria = config.PATH + "/" + session.Usuario + "/" + galery + "/images"
	os.MkdirAll(galeria, 0755)

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	var ramd = helper.GenerateRandomString(20)
	var name = galeria + "/" + group + ".jpg"
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
	helper.BuildMini(tnme, name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}
