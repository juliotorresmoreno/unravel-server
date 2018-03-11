package profile

import (
	"net/url"

	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// update actualiza los datos del perfil
func updateAll(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) (err error) {
	if legenda := data.Get("legenda"); legenda != "" {
		perfil.Legenda = legenda
	}
	if descripcion := data.Get("descripcion"); descripcion != "" {
		perfil.Descripcion = descripcion
	}
	if precioHora := data.Get("precio_hora"); precioHora != "" {
		perfil.PrecioHora = precioHora
	}
	if perfil.Id == 0 {
		perfil.Usuario = session.Usuario
		_, err = models.Add(perfil)
	} else {
		_, err = models.Update(perfil.Id, perfil)
	}
	return err
}
