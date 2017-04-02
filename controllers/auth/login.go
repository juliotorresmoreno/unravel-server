package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/unravel-server/config"
	"github.com/unravel-server/helper"
	"github.com/unravel-server/models"
)

// Login aqui es donde nos autenticamos
func Login(w http.ResponseWriter, r *http.Request) {
	var usuario = r.PostFormValue("usuario")
	var passwd = r.PostFormValue("passwd")
	var respuesta []byte
	users := make([]models.User, 0)
	orm := models.GetXORM()
	defer orm.Close()
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
