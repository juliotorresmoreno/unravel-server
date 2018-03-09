package auth

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
)

// Registrar aca es donde registramos los usuarios en bd
func Registrar(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	data := helper.GetPostParams(r)
	user.Nombres = data.Get("nombres")
	user.Apellidos = data.Get("apellidos")
	user.FullName = user.Nombres + " " + user.Apellidos
	user.Usuario = data.Get("usuario")
	user.Email = data.Get("email")
	user.Passwd = data.Get("passwd")
	user.Tipo = "Usuario"

	if _, err := user.Add(); err != nil {
		helper.Cors(w, r)
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
	helper.Cors(w, r)
	w.WriteHeader(http.StatusCreated)
	w.Write(respuesta)
}
