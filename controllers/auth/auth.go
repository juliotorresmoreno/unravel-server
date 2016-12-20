package auth

import (
	"../../models"
	"net/http"
	"encoding/json"
	"../responses"
	"../../helper"
	"../../config"
	"time"
)

func Session(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	cache := models.GetCache()
	_token := helper.GetCookie(r, "token")
	session := cache.Get(_token)
	if session.Err() != nil {
		respuesta, _ := json.Marshal(responses.Error{
			Success: false,
			Error: session.Err().Error(),
		})
		w.Write([]byte(respuesta))
	}
	usuario, _:= session.Result()
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ?", usuario).Find(&users)
	if err != nil {
		respuesta, _ := json.Marshal(responses.Error{Success:false,Error:err.Error()})
		w.Write(respuesta)
		return
	}
	respuesta, _ := json.Marshal(responses.Login{
		Success: true,
		Session: responses.Session{
			Usuario: users[0].Usuario,
			Nombres: users[0].Nombres,
			Apellidos: users[0].Apellidos,
			Token: _token,
		},
	})
	w.Write(respuesta)
}

func autenticate(user *models.User) (string, responses.Login) {
	_token := helper.GenerateRandomString(100)
	cache := models.GetCache()
	cache.Set(string(_token), user.Usuario, time.Duration(config.SESSION_DURATION) * time.Second)

	respuesta := responses.Login{
		Success: true,
		Session: responses.Session{
			Usuario: user.Usuario,
			Nombres: user.Nombres,
			Apellidos: user.Apellidos,
			Token: _token,
		},
	}

	return _token, respuesta
}

func Login(w http.ResponseWriter, r *http.Request)  {
	var usuario string = r.PostFormValue("usuario")
	var passwd string = r.PostFormValue("passwd")
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	err := orm.Where("Usuario = ?", usuario).Find(&users)
	w.Header().Set("Content-Type", "application/json")

	if err == nil && len(users) > 0 && helper.IsValid(users[0].Passwd, passwd) {
		_token, _session := autenticate(&users[0])
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			MaxAge: config.SESSION_DURATION,
			Name: "token",
			Value: _token,
			Path: "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ = json.Marshal(_session)
		w.Write(respuesta)
		return
	}
	w.WriteHeader(http.StatusOK)

	if err != nil {
		respuesta, _ = json.Marshal(responses.Error{Success:false,Error:err.Error()})
		w.Write(respuesta)
		return
	}
	respuesta, _ = json.Marshal(responses.Error{Success:false,Error:"Usuario o contraseña invalido"})
	w.Write(respuesta)
}

func Registrar(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	if r.PostFormValue("Passwd") != r.PostFormValue("PasswdConfirm") {
		respuesta, _ := json.Marshal(responses.Error{
			Success:false,
			Error:"Debe validar la contraseña.",
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respuesta, _ := json.Marshal(responses.Error{
			Success:false,
			Error:err.Error(),
		})
		w.Write(respuesta)
	} else {
		_token, _session := autenticate(&user)
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			MaxAge: config.SESSION_DURATION,
			Name: "token",
			Value: _token,
			Path: "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ := json.Marshal(_session)
		w.Write(respuesta)
	}
}