package profile

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../../helper"
	"../../models"
	"../../ws"
	"../responses"
)

func updateProfile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var user models.User
	user.Nombres = r.PostFormValue("nombres")
	user.Apellidos = r.PostFormValue("apellidos")
	user.Usuario = session.Usuario
	if user.Nombres != "" && user.Apellidos != "" {
		if _, err := user.Update(); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			var respuesta, _ = json.Marshal(responses.Error{
				Success: false,
				Error:   err.Error(),
			})
			w.Write(respuesta)
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func getProfile(session *models.User) models.Profile {
	var perfil = make([]models.Profile, 0)
	var orm = models.GetXORM()
	var err = orm.Where("Usuario = ?", session.Usuario).Find(&perfil)
	if err != nil {
		return models.Profile{}
	}
	if len(perfil) == 1 {
		return perfil[0]
	}
	return models.Profile{}
}

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

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var perfil models.Profile = getProfile(session)
	//time.Sleep(3 * time.Second)
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

// Profile consulta de perfil.
func Profile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	var perfil = make([]models.Profile, 0)
	var orm = models.GetXORM()
	if vars["user"] != "" {
		usuario = vars["user"]
	} else {
		usuario = session.Usuario
	}
	if err := orm.Where("Usuario = ?", usuario).Find(&perfil); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(perfil) == 1 {
		var jsonData []byte
		if usuario != session.Usuario {
			estado := map[int8]string {
				-1: "Desconocidos",
				models.EstadoSolicitado: "Solicitado",
				models.EstadoAceptado: "Amigos",
			}[models.IsFriend(session.Usuario, perfil[0].Usuario)]
			jsonData, _ = json.Marshal(map[string]interface{} {
					"success": true,
					"data": truncar(perfil[0], estado),
				})
		} else {
			jsonData, _ = json.Marshal(map[string]interface{} {
				"success": true,
				"data": perfil[0],
			})
		}
		w.Write(jsonData)
		return
	}
	w.Write([]byte("{}"))
}

func truncar(p models.Profile, relacion string) map[string]string {
	var r = make(map[string]string)

	r["estado"] = relacion

	if helper.PuedoVer(relacion, p.PermisoEmail) {
		r["email"] = p.Email
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoDia) {
		r["nacimiento_mes"] = p.NacimientoDia
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoAno) {
		r["nacimiento_ano"] = p.NacimientoAno
	}
	if helper.PuedoVer(relacion, p.PermisoSexo) {
		r["sexo"] = p.Sexo
	}

	if helper.PuedoVer(relacion, p.PermisoNacimientoPais) {
		r["nacimiento_pais"] = p.NacimientoPais
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoCiudad) {
		r["nacimiento_ciudad"] = p.NacimientoCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaPais) {
		r["residencia_pais"] = p.ResidenciaPais
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaCiudad) {
		r["residencia_ciudad"] = p.ResidenciaCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoDireccion) {
		r["direccion"] = p.Direccion
	}
	if helper.PuedoVer(relacion, p.PermisoTelefono) {
		r["telefono"] = p.Telefono
	}
	if helper.PuedoVer(relacion, p.PermisoCelular) {
		r["celular"] = p.Celular
	}

	if helper.PuedoVer(relacion, p.PermisoPersonalidad) {
		r["personalidad"] = p.Personalidad
	}
	if helper.PuedoVer(relacion, p.PermisoIntereses) {
		r["intereses"] = p.Intereses
	}
	if helper.PuedoVer(relacion, p.PermisoSeries) {
		r["series"] = p.Series
	}
	if helper.PuedoVer(relacion, p.PermisoMusica) {
		r["musica"] = p.Musica
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasReligiosas) {
		r["creencias_religiosas"] = p.CreenciasReligiosas
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasPoliticas) {
		r["creencias_politicas"] = p.CreenciasPoliticas
	}

	return r
}
