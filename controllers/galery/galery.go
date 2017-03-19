package galery

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"regexp"

	"../../config"
	"../../helper"
	"../../models"
	"../../ws"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
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

// Save crea y actualiza la galeria
func Save(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var ID = strings.Trim(r.PostFormValue("ID"), " ")
	var nombre = strings.Trim(r.PostFormValue("nombre"), " ")
	var permiso = r.PostFormValue("permiso")
	var descripcion = r.PostFormValue("descripcion")

	w.Header().Set("Content-Type", "application/json")
	if !helper.IsValidPermision(permiso) || !govalidator.IsAlphanumeric(nombre) {
		helper.DespacharError(w, errors.New("El nombre es invalido"), http.StatusNotAcceptable)
		return
	}

	var galeria = config.PATH + "/" + session.Usuario + "/" + strings.Trim(nombre, "\n")
	if ID != "" {
		var galeriaOld = config.PATH + "/" + session.Usuario + "/" + strings.Trim(ID, "\n")
		if _, err := os.Stat(galeriaOld); err != nil {
			helper.DespacharError(w, err, http.StatusInternalServerError)
			return
		}
		os.Rename(galeriaOld, galeria)
	} else {
		if _, err := os.Stat(galeria); err != nil {
			if err = os.MkdirAll(galeria, 0755); err != nil {
				helper.DespacharError(w, err, http.StatusInternalServerError)
				return
			}
		}
	}

	p, _ := os.Create(galeria + "/permiso")
	defer p.Close()
	p.Write([]byte(permiso))

	d, _ := os.Create(galeria + "/descripcion")
	defer d.Close()
	d.Write([]byte(descripcion))

	w.WriteHeader(http.StatusCreated)
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"galeria": nombre,
	})
	w.Write(respuesta)
}

func describeGaleria(usuario, galeria string) (string, string, error) {
	var path = config.PATH + "/" + usuario
	permiso, err := ioutil.ReadFile(path + "/" + galeria + "/permiso")
	if err != nil {
		return string(permiso), "", err
	}
	descripcion, err := ioutil.ReadFile(path + "/" + galeria + "/descripcion")
	if err != nil {
		return string(permiso), string(descripcion), err
	}
	return string(permiso), string(descripcion), nil
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
		permiso, descripcion, err := describeGaleria(usuario, files[i].Name())
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
	var files, _ = ioutil.ReadDir(path + "/" + galeria + "/images")
	var length = len(files)
	var imagenes = make([]string, length)
	for i := 0; i < length; i++ {
		imagenes[i] = strings.Trim(files[i].Name(), "\n")
	}
	return imagenes
}

var nombreValido, _ = regexp.Compile("^[A-Za-z0-9\\.]+$")
var galeriaValida, _ = regexp.Compile("^[A-Za-z0-9]+$")

// EliminarImagen imagenes de la galerias existente
func EliminarImagen(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var galeria = r.PostFormValue("galery")
	var imagen = r.PostFormValue("image")
	var usuario = session.Usuario
	if !nombreValido.MatchString(imagen) || !galeriaValida.MatchString(galeria) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"success\": false}"))
		return
	}
	var path = config.PATH + "/" + usuario + "/" + galeria + "/images/" + imagen
	var pathm = config.PATH + "/" + usuario + "/" + galeria + "/mini/" + imagen
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		os.Remove(path)
		os.Remove(pathm)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
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

// DescribeGaleria ver imagen
func DescribeGaleria(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var galeria = vars["galery"]
	var usuario string
	if vars["usuario"] != "" {
		usuario = vars["usuario"]
	} else {
		usuario = session.Usuario
	}
	var permiso, descripcion, err = describeGaleria(usuario, galeria)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"ID":          galeria,
			"nombre":      galeria,
			"permiso":     permiso,
			"descripcion": descripcion,
		},
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

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
