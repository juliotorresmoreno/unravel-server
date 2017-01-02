package profile

import (
	"encoding/json"
	"net/http"

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

func updateEmail(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.WriteHeader(http.StatusNoContent)
}

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	if r.PostFormValue("nombres") != "" && r.PostFormValue("apellidos") != "" {
		updateProfile(w, r, session, hub)
		return
	}
	if r.PostFormValue("permiso_email") != "" && r.PostFormValue("email") != "" {
		updateEmail(w, r, session, hub)
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
