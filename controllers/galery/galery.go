package galery

import (
	"net/http"

	"../../helper"
	"../../models"
	"../../ws"
	"../../config"
	"os"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var nombre = strings.Trim(r.PostFormValue("nombre"), " ")
	var permiso = r.PostFormValue("permiso")
	var descripcion = r.PostFormValue("descripcion")

	w.Header().Set("Content-Type", "application/json")
	if !helper.IsValidPermision(permiso) || nombre == "" || strings.Contains(nombre, ".") {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("{\"success\": false}"))
		return
	}

	var galeria = config.PATH + "/" + session.Usuario + "/" + nombre
	if _, err:= os.Stat(galeria); err != nil {
		if err = os.MkdirAll(galeria, 0755); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("{\"success\": false}"))
			return
		}
	}

	p, _ := os.Create(galeria + "/permiso")
	p.Write([]byte(permiso))
	defer p.Close()

	d, _ := os.Create(galeria + "/descripcion")
	d.Write([]byte(descripcion))
	defer d.Close()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\": true}"))
}