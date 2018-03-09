package profile

import (
	"net/http"
	"net/url"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func updateEmail(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	if !helper.IsValidPermision(data.Get("permiso_email")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var err error
	perfil.Email = data.Get("email")
	perfil.PermisoEmail = data.Get("permiso_email")
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
	w.WriteHeader(http.StatusNoContent)
}

//
func updateNacimientoMes(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_nacimiento_dia")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoDia = data.Get("nacimiento_dia")
	perfil.NacimientoMes = data.Get("nacimiento_mes")
	perfil.PermisoNacimientoDia = data.Get("permiso_nacimiento_dia")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateNacimientoAno(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_nacimiento_ano")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoAno = data.Get("nacimiento_ano")
	perfil.PermisoNacimientoAno = data.Get("permiso_nacimiento_ano")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateNacimientoSexo(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_sexo")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Sexo = data.Get("sexo")
	perfil.PermisoSexo = data.Get("permiso_sexo")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateNacimientoPais(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_nacimiento_pais")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoPais = data.Get("nacimiento_pais")
	perfil.PermisoNacimientoPais = data.Get("permiso_nacimiento_pais")
	if perfil.Id == 0 {
		perfil.Usuario = session.Usuario
		_, err = models.Add(perfil)
	} else {
		_, err = models.Update(perfil.Id, perfil)
	}
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateNacimientoCiudad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_nacimiento_ciudad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoCiudad = data.Get("nacimiento_ciudad")
	perfil.PermisoNacimientoCiudad = data.Get("permiso_nacimiento_ciudad")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateResidenciaPais(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_residencia_pais")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.ResidenciaPais = data.Get("residencia_pais")
	perfil.PermisoResidenciaPais = data.Get("permiso_residencia_pais")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateResidenciaCiudad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_residencia_ciudad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.ResidenciaCiudad = data.Get("residencia_ciudad")
	perfil.PermisoResidenciaCiudad = data.Get("permiso_residencia_ciudad")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateDireccion(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_direccion")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Direccion = data.Get("direccion")
	perfil.PermisoDireccion = data.Get("permiso_direccion")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateTelefono(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_telefono")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Telefono = data.Get("telefono")
	perfil.PermisoTelefono = data.Get("permiso_telefono")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateCelular(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_celular")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Celular = data.Get("celular")
	perfil.PermisoCelular = data.Get("permiso_celular")
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
	w.WriteHeader(http.StatusNoContent)
}

func updatePersonalidad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_personalidad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Personalidad = data.Get("personalidad")
	perfil.PermisoPersonalidad = data.Get("permiso_personalidad")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateIntereses(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_intereses")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Intereses = data.Get("intereses")
	perfil.PermisoIntereses = data.Get("permiso_intereses")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateMusica(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_musica")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Musica = data.Get("musica")
	perfil.PermisoMusica = data.Get("permiso_musica")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateSeries(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_series")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Series = data.Get("series")
	perfil.PermisoSeries = data.Get("permiso_series")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateCreenciasReligiosas(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_creencias_religiosas")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.CreenciasReligiosas = data.Get("creencias_religiosas")
	perfil.PermisoCreenciasReligiosas = data.Get("permiso_creencias_religiosas")
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
	w.WriteHeader(http.StatusNoContent)
}

func updateCreenciasPoliticas(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile, data url.Values) {
	var err error

	if !helper.IsValidPermision(data.Get("permiso_creencias_politicas")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.CreenciasPoliticas = data.Get("creencias_politicas")
	perfil.PermisoCreenciasPoliticas = data.Get("permiso_creencias_politicas")
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
	w.WriteHeader(http.StatusNoContent)
}
