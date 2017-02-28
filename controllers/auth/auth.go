package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"../../config"
	"../../helper"
	"../../models"
	"../../ws"
)

// Logout cerrar session
func Logout(w http.ResponseWriter, r *http.Request) {
	var cache = models.GetCache()
	var token string = helper.GetToken(r)
	if token != "" {
		cache.Del(token)
		http.SetCookie(w, &http.Cookie{
			MaxAge:   0,
			Secure:   false,
			HttpOnly: true,
			Name:     "token",
			Value:    "",
			Path:     "/",
			Expires:  time.Now(),
		})
	}
	helper.Cors(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":true}"))
}

// Session obtiene la session actual del usuario logueado
func Session(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	if session == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"success\": false}"))
		return
	}
	var _token string = helper.GetCookie(r, "token")
	if _token == "" {
		_token = r.URL.Query().Get("token")
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"session": map[string]string{
			"usuario":   session.Usuario,
			"nombres":   session.Nombres,
			"fullname":  session.FullName,
			"apellidos": session.Apellidos,
			"token":     _token,
		},
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

func autenticate(user *models.User) (string, map[string]interface{}) {
	_token := helper.GenerateRandomString(100)
	cache := models.GetCache()
	cache.Set(string(_token), user.Usuario, time.Duration(config.SESSION_DURATION)*time.Second)

	respuesta := map[string]interface{}{
		"success": true,
		"session": map[string]string{
			"usuario":   user.Usuario,
			"fullname":  user.FullName,
			"nombres":   user.Nombres,
			"apellidos": user.Apellidos,
			"token":     _token,
		},
	}

	return _token, respuesta
}

// Login aqui es donde nos autenticamos
func Login(w http.ResponseWriter, r *http.Request) {
	var usuario = r.PostFormValue("usuario")
	var passwd = r.PostFormValue("passwd")
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ?", usuario).Find(&users)
	w.Header().Set("Content-Type", "application/json")
	helper.Cors(w, r)

	if err == nil && len(users) > 0 && helper.IsValid(users[0].Passwd, passwd) {
		_token, _session := autenticate(&users[0])
		http.SetCookie(w, &http.Cookie{
			MaxAge:   config.SESSION_DURATION,
			Secure:   false,
			HttpOnly: true,
			Name:     "token",
			Value:    _token,
			Path:     "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ = json.Marshal(_session)
		w.Write(respuesta)
		return
	}

	if err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}

	helper.DespacharError(w, errors.New("Usuario o contraseña invalido"), http.StatusUnauthorized)
}

// Registrar aca es donde registramos los usuarios en bd
func Registrar(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	if r.PostFormValue("passwd") != "" && r.PostFormValue("passwd") != r.PostFormValue("passwdConfirm") {
		helper.DespacharError(w, errors.New("Passwd: Debe validar la contraseña"), http.StatusNotAcceptable)
		return
	}

	user.Nombres = r.PostFormValue("nombres")
	user.Apellidos = r.PostFormValue("apellidos")
	user.FullName = user.Nombres + " " + user.Apellidos
	user.Usuario = r.PostFormValue("usuario")
	user.Email = r.PostFormValue("email")
	user.Passwd = r.PostFormValue("passwd")

	if _, err := user.Add(); err != nil {
		helper.DespacharError(w, err, http.StatusNotAcceptable)
		return
	}
	var _token, _session = autenticate(&user)
	var respuesta, _ = json.Marshal(_session)
	http.SetCookie(w, &http.Cookie{
		MaxAge:   config.SESSION_DURATION,
		HttpOnly: true,
		Secure:   false,
		Name:     "token",
		Value:    _token,
		Path:     "/",
	})
	w.WriteHeader(http.StatusCreated)
	w.Write(respuesta)
}
