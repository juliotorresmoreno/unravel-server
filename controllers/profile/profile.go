package profile

import (
	"encoding/json"
	"net/http"

	"time"

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

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var perfil models.Profile = getProfile(session)
	time.Sleep(3 * time.Second)
	if r.PostFormValue("nombres") != "" && r.PostFormValue("apellidos") != "" {
		updateProfile(w, r, session, hub)
		return
	}
	if r.PostFormValue("permiso_email") != "" && r.PostFormValue("email") != "" {
		updateEmail(w, r, session, hub, perfil)
		return
	}
	if r.PostFormValue("permiso_nacimiento_dia") != "" && r.PostFormValue("nacimiento_dia") != "" &&
		r.PostFormValue("nacimiento_mes") != "" {
		updateNacimientoMes(w, r, session, hub, perfil)
		return
	}
}

// Profile consulta de perfil.
func Profile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var perfil = make([]models.Profile, 0)
	var orm = models.GetXORM()
	var err = orm.Where("Usuario = ?", session.Usuario).Find(&perfil)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(perfil) == 1 {
		json, _ := json.Marshal(perfil[0])
		w.Write(json)
		return
	}
	w.Write([]byte("{\"success\": true, \"data\": {}}"))
}
