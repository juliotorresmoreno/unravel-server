package galery

import (
	"net/http"
	"os"

	"../../config"
	"../../models"
	"../../ws"
)

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
