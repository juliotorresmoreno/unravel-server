package galery

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/unravel-server/config"
	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
	"github.com/unravel-server/ws"
)

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
