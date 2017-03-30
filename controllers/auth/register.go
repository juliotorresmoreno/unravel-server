package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/unravel-server/config"
	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
)

// Registrar aca es donde registramos los usuarios en bd
func Registrar(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	if r.PostFormValue("passwd") != "" && r.PostFormValue("passwd") != r.PostFormValue("passwdConfirm") {
		helper.Cors(w, r)
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
