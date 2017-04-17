package profile

import (
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func updateEmail(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	if !helper.IsValidPermision(r.PostFormValue("permiso_email")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var err error
	perfil.Email = r.PostFormValue("email")
	perfil.PermisoEmail = r.PostFormValue("permiso_email")
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
func updateNacimientoMes(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_nacimiento_dia")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoDia = r.PostFormValue("nacimiento_dia")
	perfil.NacimientoMes = r.PostFormValue("nacimiento_mes")
	perfil.PermisoNacimientoDia = r.PostFormValue("permiso_nacimiento_dia")
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

func updateNacimientoAno(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_nacimiento_ano")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoAno = r.PostFormValue("nacimiento_ano")
	perfil.PermisoNacimientoAno = r.PostFormValue("permiso_nacimiento_ano")
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

func updateNacimientoSexo(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_sexo")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Sexo = r.PostFormValue("sexo")
	perfil.PermisoSexo = r.PostFormValue("permiso_sexo")
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

func updateNacimientoPais(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_nacimiento_pais")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoPais = r.PostFormValue("nacimiento_pais")
	perfil.PermisoNacimientoPais = r.PostFormValue("permiso_nacimiento_pais")
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

func updateNacimientoCiudad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_nacimiento_ciudad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.NacimientoCiudad = r.PostFormValue("nacimiento_ciudad")
	perfil.PermisoNacimientoCiudad = r.PostFormValue("permiso_nacimiento_ciudad")
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

func updateResidenciaPais(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_residencia_pais")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.ResidenciaPais = r.PostFormValue("residencia_pais")
	perfil.PermisoResidenciaPais = r.PostFormValue("permiso_residencia_pais")
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

func updateResidenciaCiudad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_residencia_ciudad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.ResidenciaCiudad = r.PostFormValue("residencia_ciudad")
	perfil.PermisoResidenciaCiudad = r.PostFormValue("permiso_residencia_ciudad")
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

func updateDireccion(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_direccion")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Direccion = r.PostFormValue("direccion")
	perfil.PermisoDireccion = r.PostFormValue("permiso_direccion")
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

func updateTelefono(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_telefono")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Telefono = r.PostFormValue("telefono")
	perfil.PermisoTelefono = r.PostFormValue("permiso_telefono")
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

func updateCelular(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_celular")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Celular = r.PostFormValue("celular")
	perfil.PermisoCelular = r.PostFormValue("permiso_celular")
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

func updatePersonalidad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_personalidad")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Personalidad = r.PostFormValue("personalidad")
	perfil.PermisoPersonalidad = r.PostFormValue("permiso_personalidad")
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

func updateIntereses(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_intereses")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Intereses = r.PostFormValue("intereses")
	perfil.PermisoIntereses = r.PostFormValue("permiso_intereses")
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

func updateMusica(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_musica")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Musica = r.PostFormValue("musica")
	perfil.PermisoMusica = r.PostFormValue("permiso_musica")
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

func updateSeries(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_series")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.Series = r.PostFormValue("series")
	perfil.PermisoSeries = r.PostFormValue("permiso_series")
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

func updateCreenciasReligiosas(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_creencias_religiosas")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.CreenciasReligiosas = r.PostFormValue("creencias_religiosas")
	perfil.PermisoCreenciasReligiosas = r.PostFormValue("permiso_creencias_religiosas")
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

func updateCreenciasPoliticas(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil models.Profile) {
	var err error
	if !helper.IsValidPermision(r.PostFormValue("permiso_creencias_politicas")) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	perfil.CreenciasPoliticas = r.PostFormValue("creencias_politicas")
	perfil.PermisoCreenciasPoliticas = r.PostFormValue("permiso_creencias_politicas")
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
