package galery

import (
	"net/http"

	"../../helper"
	"../../models"
	"../../ws"
	"../../config"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
)

// Upload sube las imagenes
func Upload(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("{\"success\": true}"))

	var galeria = r.PostFormValue("galery")
	var name = r.PostFormValue("name")
	var file = r.PostFormValue("file")

	println(name, galeria)
	println(len(file))
	w.Write([]byte(file))
}

// Create crea la nueva galeria
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

// Listar lista las galerias existentes
func Listar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = config.PATH + "/" + session.Usuario
	var files, _ = ioutil.ReadDir(usuario)
	var length = len(files)
	var list = make([]string, length)
	for i := 0; i < length; i++ {
		list[i] = files[i].Name()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{} {
		"success": true,
		"data": list,
	})
	w.Write([]byte(respuesta))
}