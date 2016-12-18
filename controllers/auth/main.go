package auth

import (
	"../../models"
	"net/http"
	"encoding/json"
	"../responses"
	"../../services"
	"../../config"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	var usuario string = r.PostFormValue("usuario")
	var passwd string = r.PostFormValue("passwd")
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	cache := models.GetCache()
	err := orm.Where("Usuario = ?", usuario).Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		respuesta, _ = json.Marshal(responses.Error{Success:false,Error:err.Error()})
		w.Write(respuesta)
		return
	}

	if len(users) > 0 && services.IsValid(users[0].Passwd, passwd) {
		_token := services.GenerateRandomString(100)
		respuesta, _ = json.Marshal(responses.Login{
			Success: true,
			Session: responses.Session{
				Usuario: users[0].Usuario,
				Nombres: users[0].Nombres,
				Apellidos: users[0].Apellidos,
				Token: _token,
			}})
		cache.Set(string(_token), respuesta, config.SESSION_DURATION)
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
		w.WriteHeader(http.StatusOK)
	}
}