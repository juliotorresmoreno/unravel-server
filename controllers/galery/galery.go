package galery

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"fmt"

	"../../config"
	"../../helper"
	"../../models"
	"../../ws"
)

// Upload sube las imagenes
func Upload(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var galery = strings.Trim(r.FormValue("galery"), " ")
	var galeria = config.PATH + "/" + session.Usuario + "/" + galery
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	var _name = strings.TrimLeft(header.Filename, ".")
	var index = 1
	var name = _name
	for {
		if _, err := os.Stat(galeria + "/" + name); err != nil {
			break
		}
		name = fmt.Sprintln(index, _name)
		index++
	}
	println(name)

	save, _ := os.Create(galeria + "/" + name)
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
	if _, err := os.Stat(galeria); err != nil {
		if err = os.MkdirAll(galeria, 0755); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("{\"success\": false}"))
			return
		}
	}

	p, _ := os.Create(galeria + "/permiso")
	defer p.Close()
	p.Write([]byte(permiso))

	d, _ := os.Create(galeria + "/descripcion")
	defer d.Close()
	d.Write([]byte(descripcion))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\": true}"))
}

// ListarGalerias lista las galerias existentes
func ListarGalerias(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = config.PATH + "/" + session.Usuario
	var files, _ = ioutil.ReadDir(usuario)
	var length = len(files)
	var list = make([]interface{}, length)
	for i := 0; i < length; i++ {
		permiso, _ := ioutil.ReadFile(usuario + "/" + files[i].Name() + "/permiso")
		descripcion, _ := ioutil.ReadFile(usuario + "/" + files[i].Name() + "/descripcion")
		list[i] = map[string]interface{}{
			"name":        files[i].Name(),
			"permiso":     string(permiso),
			"descripcion": string(descripcion),
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    list,
	})
	w.Write([]byte(respuesta))
}

// Listar imagenes de la galerias existente
func ListarImagenes(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var path = config.PATH + "/" + usuario
	var files, _ = ioutil.ReadDir(path + "/" + galeria)
	var length = len(files)
	var list = make([]string, 0)
	for i := 0; i < length; i++ {
		if files[i].Name() != "descripcion" && files[i].Name() != "permiso" {
			list = append(list, strings.Trim(files[i].Name(), "\n"))
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    list,
	})
	w.Write([]byte(respuesta))
}
