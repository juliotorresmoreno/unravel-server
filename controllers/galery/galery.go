package galery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/middlewares"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func NewRouter(hub *ws.Hub) http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/", middlewares.Protect(ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/", middlewares.Protect(Save, hub, true)).Methods("POST")
	mux.HandleFunc("/delete", middlewares.Protect(EliminarImagen, hub, true)).Methods("POST", "DELETE")
	mux.HandleFunc("/fotoPerfil", middlewares.Protect(GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/upload", middlewares.Protect(Upload, hub, true)).Methods("POST")
	mux.HandleFunc("/fotoPerfil", middlewares.Protect(SetFotoPerfil, hub, true)).Methods("POST")
	mux.HandleFunc("/fotoPerfil/{usuario}", middlewares.Protect(GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/{galery}/describe", middlewares.Protect(DescribeGaleria, hub, true)).Methods("GET")
	mux.HandleFunc("/{galery}/preview", middlewares.Protect(ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/{galery}/{imagen}", middlewares.Protect(ViewImagen, hub, true)).Methods("GET")
	mux.HandleFunc("/{galery}", middlewares.Protect(ListarImagenes, hub, true)).Methods("GET")

	return mux
}

func NewUserRouter(hub *ws.Hub) http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/{usuario}/galery", middlewares.Protect(ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/{usuario}/galery/fotoPerfil", middlewares.Protect(GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/{usuario}/galery/{galery}", middlewares.Protect(ListarImagenes, hub, true)).Methods("GET")
	mux.HandleFunc("/{usuario}/galery/{galery}/preview", middlewares.Protect(ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/{usuario}/galery/{galery}/{imagen}", middlewares.Protect(ViewImagen, hub, true)).Methods("GET")

	return mux
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

type image struct {
	Src      string    `json:"src"`
	Modified time.Time `json:"modified"`
}

func listarImagenes(usuario, galeria string) []image {
	var path = config.PATH + "/" + usuario
	var files, _ = ioutil.ReadDir(path + "/" + galeria + "/images")
	var length = len(files)
	var imagenes = make([]image, length)

	for i := 0; i < length; i++ {
		imagenes[i] = image{
			Src:      strings.Trim(files[i].Name(), "\n"),
			Modified: files[i].ModTime(),
		}
	}

	return imagenes
}

var nombreValido, _ = regexp.Compile("^[A-Za-z0-9\\.]+$")
var galeriaValida, _ = regexp.Compile("^[A-Za-z0-9]+$")

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
