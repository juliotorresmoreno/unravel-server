package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

var templates = template.Must(template.ParseGlob("templates/*"))

// Recovery recuperacion de contraseña
func Recovery(w http.ResponseWriter, r *http.Request) {
	helper.Cors(w, r)
	data := helper.GetPostParams(r)
	email := data.Get("email")
	user, err := buscarUsuario(email)
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	var orm = db.GetXORM()
	defer orm.Close()
	user.Recovery = helper.GenerateRandomString(100)
	_, err = orm.Id(user.Id).Cols("recovery").Update(user)
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	to := []string{email}
	msg := render(email, user.Recovery)
	server := config.SMTP_HOST + ":" + config.SMTP_PORT
	err = smtp.SendMail(server, nil, "Recuperacion de contraseña <recovery@onnasoft.com>", to, msg)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}

func buscarUsuario(email string) (*models.User, error) {
	users := make([]models.User, 0)
	orm := db.GetXORM()
	defer orm.Close()
	err := orm.Where("Email = ?", email).Find(&users)
	if err != nil {
		return &models.User{}, err
	}
	if len(users) == 0 {
		return &models.User{}, errors.New("El usuario no existe")
	}
	return &users[0], nil
}

func buscarID(id string) (*models.User, error) {
	users := make([]models.User, 0)
	orm := db.GetXORM()
	defer orm.Close()
	err := orm.Where("Recovery = ?", id).Find(&users)
	if err != nil {
		return &models.User{}, err
	}
	if len(users) == 0 {
		return &models.User{}, errors.New("El token ha caducado")
	}
	return &users[0], nil
}

//Password recupera la contraseña del usuario
func Password(w http.ResponseWriter, r *http.Request) {
	helper.Cors(w, r)
	data := helper.GetPostParams(r)
	id := data.Get("id")
	password := data.Get("passwd")
	cpassword := data.Get("passwdConfirm")

	if password != cpassword {
		helper.DespacharError(w, errors.New("Passwd: Debe validar la contraseña"), http.StatusNotAcceptable)
		return
	}
	user, err := buscarID(id)
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotFound)
		return
	}
	if govalidator.IsAlphanumeric(password) == false {
		helper.DespacharError(w, errors.New("No es una contraseña valida"), http.StatusBadRequest)
		return
	}
	orm := db.GetXORM()
	defer orm.Close()
	user.Passwd = password
	user.Recovery = ""
	_, err = orm.Id(user.Id).Cols("passwd", "recovery").Update(user)
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	_token, _session := autenticate(user)
	http.SetCookie(w, &http.Cookie{
		MaxAge:   config.SESSION_DURATION,
		Secure:   false,
		HttpOnly: true,
		Name:     "token",
		Value:    _token,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(_session)
	w.Write(respuesta)
	return
}

//PasswordChange recupera la contraseña del usuario
func PasswordChange(w http.ResponseWriter, r *http.Request, user *models.User, hub *ws.Hub) {
	helper.Cors(w, r)
	data := helper.GetPostParams(r)
	password := data.Get("passwd")
	passwordNew := data.Get("passwdNew")

	if !helper.IsValid(user.Passwd, password) {
		helper.DespacharError(w, errors.New("No es la contraseña valida"), http.StatusBadRequest)
		return
	}

	if govalidator.IsAlphanumeric(password) == false {
		helper.DespacharError(w, errors.New("No es una nueva contraseña valida"), http.StatusBadRequest)
		return
	}
	orm := db.GetXORM()
	defer orm.Close()
	user.Passwd = passwordNew
	user.Recovery = ""

	if ok, err := govalidator.ValidateStruct(user); !ok {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}

	user.Passwd = helper.Encript(passwordNew)

	_, err := orm.Id(user.Id).Cols("passwd", "recovery").Update(user)
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	_token, _session := autenticate(user)
	http.SetCookie(w, &http.Cookie{
		MaxAge:   config.SESSION_DURATION,
		Secure:   false,
		HttpOnly: true,
		Name:     "token",
		Value:    _token,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(_session)
	w.Write(respuesta)
	return
}
