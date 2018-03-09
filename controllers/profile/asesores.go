package profile

import (
	"net/http"
	"net/url"

	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// update actualiza los datos del perfil
func updateAll(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	update := false
	if legenda := data.Get("legenda"); legenda != "" {
		perfil.Legenda = legenda
		update = true
	}
	var err error
	if update {
		if perfil.Id == 0 {
			perfil.Usuario = session.Usuario
			_, err = models.Add(perfil)
		} else {
			_, err = models.Update(perfil.Id, perfil)
		}
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
