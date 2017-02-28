package profile

import "encoding/json"
import "net/http"

import "github.com/gorilla/mux"

import "../../helper"
import "../../models"
import "../../ws"

func updateProfile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var user models.User
	user.Nombres = r.PostFormValue("nombres")
	user.Apellidos = r.PostFormValue("apellidos")
	user.FullName = user.Nombres + " " + user.Apellidos
	user.Usuario = session.Usuario
	if user.Nombres != "" && user.Apellidos != "" {
		if _, err := user.Update(); err != nil {
			helper.DespacharError(w, err, http.StatusNotAcceptable)
			return
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
			var estado = models.IsFriend(session.Usuario, perfil[0].Usuario)
			jsonData, _ = json.Marshal(map[string]interface{}{
				"success": true,
				"data":    truncar(perfil[0], estado),
			})
		} else {
			jsonData, _ = json.Marshal(map[string]interface{}{
				"success": true,
				"data":    perfil[0],
			})
		}
		w.Write(jsonData)
		return
	}
	w.Write([]byte("{\"success\": true, \"data\": {}}"))
}
