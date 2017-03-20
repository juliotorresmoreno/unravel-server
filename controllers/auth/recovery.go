package auth

import (
	"errors"
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/asaskevich/govalidator"

	"../../config"
	"../../helper"
	"../../models"
)

var templates = template.Must(template.ParseGlob("templates/*"))

// Recovery recuperacion de contraseña
func Recovery(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	user, err := buscarUsuario(email)
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	user.Recovery = helper.GenerateRandomString(100)
	_, err = models.Update(user.Id, user)
	to := []string{email}
	msg := render(email, user.Recovery)
	server := config.SMTP_HOST + ":" + config.SMTP_PORT
	err = smtp.SendMail(server, nil, "Recuperacion de contraseña <recovery@unravel.ga>", to, msg)
	if err != nil {
		println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}

func buscarUsuario(email string) (*models.User, error) {
	users := make([]models.User, 0)
	orm := models.GetXORM()
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
	orm := models.GetXORM()
	err := orm.Where("Recovery = ?", id).Find(&users)
	if err != nil {
		return &models.User{}, err
	}
	if len(users) == 0 {
		return &models.User{}, errors.New("El usuario no existe")
	}
	return &users[0], nil
}

//Password recupera la contraseña del usuario
func Password(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	password := r.PostFormValue("password")
	cpassword := r.PostFormValue("passwdConfirm")
	if password != cpassword {
		helper.DespacharError(w, errors.New("Passwd: Debe validar la contraseña"), http.StatusNotAcceptable)
		return
	}
	user, err := buscarID(id)
	if err != nil {
		helper.DespacharError(w, err, http.StatusNotFound)
		return
	}
	user.Passwd = password
	if _, err := govalidator.ValidateStruct(user); err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	if err != nil {
		helper.DespacharError(w, err, http.StatusBadRequest)
		return
	}
	_, err = models.Update(user.Id, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}
