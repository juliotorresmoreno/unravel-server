package profile

import "net/http"
import "../../models"
import "../../ws"

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var perfil models.Profile = getProfile(session)
	if r.PostFormValue("nombres") != "" && r.PostFormValue("apellidos") != "" {
		updateProfile(w, r, session, hub)
		return
	}

	println(r.PostFormValue("permiso_residencia_pais"))
	if r.PostFormValue("permiso_email") != "" && r.PostFormValue("email") != "" {
		updateEmail(w, r, session, hub, perfil)
		return
	}
	if r.PostFormValue("permiso_nacimiento_dia") != "" && r.PostFormValue("nacimiento_dia") != "" &&
		r.PostFormValue("nacimiento_mes") != "" {
		updateNacimientoMes(w, r, session, hub, perfil)
		return
	}
	if r.PostFormValue("permiso_nacimiento_ano") != "" && r.PostFormValue("nacimiento_ano") != "" {
		updateNacimientoAno(w, r, session, hub, perfil)
		return
	}
	if r.PostFormValue("permiso_sexo") != "" && r.PostFormValue("sexo") != "" {
		updateNacimientoSexo(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_nacimiento_pais") != "" && r.PostFormValue("nacimiento_pais") != "" {
		updateNacimientoPais(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_nacimiento_ciudad") != "" && r.PostFormValue("nacimiento_ciudad") != "" {
		updateNacimientoCiudad(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_residencia_pais") != "" && r.PostFormValue("residencia_pais") != "" {
		updateResidenciaPais(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_residencia_ciudad") != "" && r.PostFormValue("residencia_ciudad") != "" {
		updateResidenciaCiudad(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_direccion") != "" && r.PostFormValue("direccion") != "" {
		updateDireccion(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_telefono") != "" && r.PostFormValue("telefono") != "" {
		updateTelefono(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_celular") != "" && r.PostFormValue("celular") != "" {
		updateCelular(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_personalidad") != "" && r.PostFormValue("personalidad") != "" {
		updatePersonalidad(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_intereses") != "" && r.PostFormValue("intereses") != "" {
		updateIntereses(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_series") != "" && r.PostFormValue("series") != "" {
		updateSeries(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_musica") != "" && r.PostFormValue("musica") != "" {
		updateMusica(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_creencias_religiosas") != "" && r.PostFormValue("creencias_religiosas") != "" {
		updateCreenciasReligiosas(w, r, session, hub, perfil)
		return
	}

	if r.PostFormValue("permiso_creencias_politicas") != "" && r.PostFormValue("creencias_politicas") != "" {
		updateCreenciasPoliticas(w, r, session, hub, perfil)
		return
	}
}
