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

func Login(w http.ResponseWriter, r *http.Request)  {
	var usuario string = r.PostFormValue("usuario")
	var passwd string = r.PostFormValue("passwd")
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	cache := models.GetCache()
	err := orm.Where("Usuario = ?", usuario).Find(&users)
	w.Header().Set("Content-Type", "application/json")

	if err == nil && len(users) > 0 && helper.IsValid(users[0].Passwd, passwd) {
		_token := helper.GenerateRandomString(100)
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			MaxAge: config.SESSION_DURATION,
			Value: "token=" + _token,
			Path: "/",
		})
		w.WriteHeader(http.StatusOK)
		respuesta, _ = json.Marshal(responses.Login{
			Success: true,
			Session: responses.Session{
				Usuario: users[0].Usuario,
				Nombres: users[0].Nombres,
				Apellidos: users[0].Apellidos,
				Token: _token,
			},
		})
		cache.Set(string(_token), users[0].Usuario, time.Duration(config.SESSION_DURATION) * time.Second)
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
	w.WriteHeader(http.StatusOK)

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

	if affected, err := user.Add(); err != nil && err.Error() != "" {
		respuesta, _ := json.Marshal(responses.Error{
			Success:false,
			Error:err.Error(),
		})
		w.Write(respuesta)
	} else if affected == 0 {
		respuesta, _ := json.Marshal(responses.Error{
			Success:false,
			Error: "No se insertaron registros.",
		})
		w.Write(respuesta)
	} else {
		respuesta, _ := json.Marshal(responses.Success{
			Success:false,
			Message: "No se insertaron registros.",
		})
		w.Write(respuesta)
	}
}