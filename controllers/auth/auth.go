package auth

import (
	"../../models"
	"fmt"
	"net/http"
)

func Registrar(w http.ResponseWriter, r *http.Request)  {
	var user models.Users
	if r.PostFormValue("Passwd") != r.PostFormValue("PasswdConfirm") {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "{success: false, error: \"Debe validar la contraseña.\"}")
		return
	}

	user.Nombres = r.PostFormValue("nombres")
	user.Apellidos = r.PostFormValue("apellidos")
	user.Usuario = r.PostFormValue("usuario")
	user.Email = r.PostFormValue("email")
	user.Passwd = r.PostFormValue("passwd")

	if affected, err := user.Add(); err != nil && err.Error() != "" {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "{success: false, error: \"" + err.Error() + "\"}")
	} else if affected == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "{success: false, error: \"No se insertaron registros\"}")
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}