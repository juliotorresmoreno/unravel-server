package profile

import (
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	perfil := getProfile(session)
	data := helper.GetPostParams(r)
	if data.Get("nombres") != "" && data.Get("apellidos") != "" {
		updateProfile(w, r, session, hub, data)
		return
	}

	if data.Get("permiso_email") != "" && data.Get("email") != "" {
		updateEmail(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_nacimiento_dia") != "" && data.Get("nacimiento_dia") != "" &&
		data.Get("nacimiento_mes") != "" {
		updateNacimientoMes(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_nacimiento_ano") != "" && data.Get("nacimiento_ano") != "" {
		updateNacimientoAno(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_sexo") != "" && data.Get("sexo") != "" {
		updateNacimientoSexo(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_nacimiento_pais") != "" && data.Get("nacimiento_pais") != "" {
		updateNacimientoPais(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_nacimiento_ciudad") != "" && data.Get("nacimiento_ciudad") != "" {
		updateNacimientoCiudad(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_residencia_pais") != "" && data.Get("residencia_pais") != "" {
		updateResidenciaPais(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_residencia_ciudad") != "" && data.Get("residencia_ciudad") != "" {
		updateResidenciaCiudad(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_direccion") != "" && data.Get("direccion") != "" {
		updateDireccion(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_telefono") != "" && data.Get("telefono") != "" {
		updateTelefono(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_celular") != "" && data.Get("celular") != "" {
		updateCelular(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_personalidad") != "" && data.Get("personalidad") != "" {
		updatePersonalidad(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_intereses") != "" && data.Get("intereses") != "" {
		updateIntereses(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_series") != "" && data.Get("series") != "" {
		updateSeries(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_musica") != "" && data.Get("musica") != "" {
		updateMusica(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_creencias_religiosas") != "" && data.Get("creencias_religiosas") != "" {
		updateCreenciasReligiosas(w, r, session, hub, perfil, data)
		return
	}

	if data.Get("permiso_creencias_politicas") != "" && data.Get("creencias_politicas") != "" {
		updateCreenciasPoliticas(w, r, session, hub, perfil, data)
		return
	}

	updateAll(w, r, session, hub, perfil, data)
}
