package profile

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func updateProfile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, data url.Values) {
	user := &models.User{}
	user.Nombres = data.Get("nombres")
	user.Apellidos = data.Get("apellidos")
	user.FullName = user.Nombres + " " + user.Apellidos
	user.Usuario = session.Usuario
	if user.Nombres != "" && user.Apellidos != "" {
		if _, err := user.Update(); err != nil {
			helper.DespacharError(w, err, http.StatusNotAcceptable)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("{\"success\": true}"))
}

func getProfile(session *models.User) *models.Profile {
	var perfil = make([]models.Profile, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var err = orm.Where("Usuario = ?", session.Usuario).Find(&perfil)
	if err != nil {
		return &models.Profile{}
	}
	if len(perfil) == 1 {
		return &perfil[0]
	}
	return &models.Profile{}
}

// Profile consulta de perfil.
func Profile(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	var perfil = make([]models.Profile, 0)
	var orm = db.GetXORM()
	defer orm.Close()
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
	if usuario == session.Usuario {
		_perfil := profile{
			Profile:   perfil[0],
			Estado:    "Activo",
			Nombres:   session.Nombres,
			Apellidos: session.Apellidos,
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data":    _perfil,
		})
		return
	}

	user := models.User{}
	if _, err := orm.Where("usuario = ?", usuario).Get(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	var _profile profile
	if len(perfil) == 1 {
		estado := models.IsFriend(session.Usuario, perfil[0].Usuario)
		_profile = truncar(perfil[0], estado)
	} else {
		_profile.Profile = models.Profile{
			Usuario: user.Usuario,
		}
	}

	_profile.Nombres = user.Nombres
	_profile.Apellidos = user.Apellidos
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    _profile,
	})
}
