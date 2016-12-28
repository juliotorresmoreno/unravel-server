package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"../../config"
	"../../helper"
	"../../models"
	"../responses"
)

// Session obtiene la session actual del usuario logueado
func Session(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cache = models.GetCache()
	var _token = helper.GetCookie(r, "token")
	var session = cache.Get(_token)

	if session.Err() != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"success\":false}"))
		return
	}
	usuario, _ := session.Result()
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ?", usuario).Find(&users)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		respuesta, _ := json.Marshal(responses.Error{Success: false, Error: err.Error()})
		w.Write(respuesta)
		return
	}
	respuesta, _ := json.Marshal(responses.Login{
		Success: true,
		Session: responses.Session{
			Usuario:   users[0].Usuario,
			Nombres:   users[0].Nombres,
			Apellidos: users[0].Apellidos,
			Token:     _token,
		},
	})
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

func autenticate(user *models.User) (string, responses.Login) {
	_token := helper.GenerateRandomString(100)
	cache := models.GetCache()
	cache.Set(string(_token), user.Usuario, time.Duration(config.SESSION_DURATION)*time.Second)

	respuesta := responses.Login{
		Success: true,
		Session: responses.Session{
			Usuario:   user.Usuario,
			Nombres:   user.Nombres,
			Apellidos: user.Apellidos,
			Token:     _token,
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err == nil && len(users) > 0 && helper.IsValid(users[0].Passwd, passwd) {
		_token, _session := autenticate(&users[0])
		http.SetCookie(w, &http.Cookie{
			MaxAge: config.SESSION_DURATION,
			Secure: false,
			Name:   "token",
			Value:  _token,
			Path:   "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ = json.Marshal(_session)
		w.Write(respuesta)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		respuesta, _ = json.Marshal(responses.Error{Success: false, Error: err.Error()})
		w.Write(respuesta)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	respuesta, _ = json.Marshal(responses.Error{Success: false, Error: "Usuario o contraseña invalido"})
	w.Write(respuesta)
}

// Registrar aca es donde registramos los usuarios en bd
func Registrar(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Cache-Control")
	w.Header().Set("Content-Type", "application/json")

	if r.PostFormValue("passwd") != "" && r.PostFormValue("passwd") != r.PostFormValue("passwdConfirm") {
		w.WriteHeader(http.StatusNotAcceptable)
		respuesta, _ := json.Marshal(responses.Error{
			Success: false,
			Error:   "Passwd: Debe validar la contraseña.",
		})
		w.Write(respuesta)
		return
	}

	user.Nombres = r.PostFormValue("nombres")
	user.Apellidos = r.PostFormValue("apellidos")
	user.Usuario = r.PostFormValue("usuario")
	user.Email = r.PostFormValue("email")
	user.Passwd = r.PostFormValue("passwd")

	if _, err := user.Add(); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		var respuesta, _ = json.Marshal(responses.Error{
			Success: false,
			Error:   err.Error(),
		})
		w.Write(respuesta)
	} else {
		var _token, _session = autenticate(&user)
		var respuesta, _ = json.Marshal(_session)
		http.SetCookie(w, &http.Cookie{
			MaxAge: config.SESSION_DURATION,
			Secure: false,
			Name:   "token",
			Value:  _token,
			Path:   "/",
		})
		w.WriteHeader(http.StatusCreated)
		w.Write(respuesta)
	}
}
