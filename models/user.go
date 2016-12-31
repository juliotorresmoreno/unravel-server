package models

import (
	"time"

	"errors"

	"../helper"
	"github.com/asaskevich/govalidator"
)

// User modelo de usuario
type User struct {
	Id        uint      `xorm:"bigint not null autoincr pk"`
	Nombres   string    `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Apellidos string    `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Email     string    `xorm:"varchar(200) not null unique" valid:"required,email"`
	Usuario   string    `xorm:"varchar(100) not null unique index" valid:"required,alphanum"`
	Passwd    string    `xorm:"varchar(100) not null" valid:"required,password"`
	CreateAt  time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (u User) TableName() string {
	return "users"
}

func init() {
	orm.Sync2(new(User))
}

// Add crear nuevo usuario
func (u User) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return 0, normalize(err, u)
	}
	u.Passwd = helper.Encript(u.Passwd)

	affected, err := orm.Insert(u)
	if err != nil {
		return affected, normalize(err, u)
	}
	return affected, nil
}

// Update crear nuevo usuario
func (u User) Update() (int64, error) {
	if u.Usuario == "" {
		return 0, errors.New("Usuario no especificado")
	}
	var users = make([]User, 0)
	var err = orm.Where("Usuario = ?", u.Usuario).Find(&users)
	if err != nil {
		return 0, err
	}

	_, err = govalidator.ValidateStruct(users[0])
	if err != nil {
		return 0, normalize(err, u)
	}
	users[0].Nombres = u.Nombres
	users[0].Apellidos = u.Apellidos

	affected, err := orm.Id(users[0].Id).Cols("nombres", "apellidos").Update(users[0])
	if err != nil {
		return affected, normalize(err, u)
	}
	return affected, nil
}
