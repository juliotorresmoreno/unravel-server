package galery

import "encoding/json"
import "io"
import "io/ioutil"
import "math/rand"
import "net/http"
import "os"
import "strings"

import "github.com/gorilla/mux"

import "fmt"

import "../../config"
import "../../helper"
import "../../models"
import "../../ws"

// FotoPerfil establece la foto de perfil.
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

// FotoPerfil establece la foto de perfil.
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

	var galeria = config.PATH + "/" + session.Usuario + "/" + strings.Trim(nombre, "\n")
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
	var vars = mux.Vars(r)
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var path = config.PATH + "/" + usuario
	var files, _ = ioutil.ReadDir(path)
	var length = len(files)
	var galerias = make([]interface{}, 0)
	for i := 0; i < length; i++ {
		if files[i].Name() == "fotoPerfil" {
			continue
		}
		permiso, err := ioutil.ReadFile(path + "/" + files[i].Name() + "/permiso")
		if err != nil {
			continue
		}
		descripcion, err := ioutil.ReadFile(path + "/" + files[i].Name() + "/descripcion")
		if err != nil {
			continue
		}
		c := map[string]interface{}{
			"name":        files[i].Name(),
			"permiso":     string(permiso),
			"descripcion": string(descripcion),
		}
		galerias = append(galerias, c)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    galerias,
	})
	w.Write([]byte(respuesta))
}

func listarImagenes(usuario, galeria string) []string {
	var path = config.PATH + "/" + usuario
	var files, _ = ioutil.ReadDir(path + "/" + galeria)
	var length = len(files)
	var imagenes = make([]string, 0)
	for i := 0; i < length; i++ {
		if files[i].Name() != "descripcion" && files[i].Name() != "permiso" {
			imagenes = append(imagenes, strings.Trim(files[i].Name(), "\n"))
		}
	}
	return imagenes
}

// ListarImagenes imagenes de la galerias existente
func ListarImagenes(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var imagenes = listarImagenes(usuario, galeria)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    imagenes,
	})
	w.Write([]byte(respuesta))
}

// ViewImagen ver imagen
func ViewImagen(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var imagen = vars["imagen"]
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var path = config.PATH + "/" + usuario + "/" + galeria + "/" + imagen
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
		url = "http://" + r.Host + defecto
	} else {
		imagen = imagenes[rand.Intn(length)]
		auth = "?token=" + token
		url = "http://" + r.Host + "/api/v1/" + usuario + "/galery/" + galeria + "/" + imagen + auth
	}
	http.Redirect(w, r, url, http.StatusFound)
}
