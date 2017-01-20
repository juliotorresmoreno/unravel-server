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
func (selft User) TableName() string {
	return "users"
}

func init() {
	orm.Sync2(new(User))
}

// Add crear nuevo usuario
func (self User) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(self)
	if err != nil {
		return 0, normalize(err, self)
	}
	self.Passwd = helper.Encript(self.Passwd)

	affected, err := orm.Insert(self)
	if err != nil {
		return affected, normalize(err, self)
	}
	return affected, nil
}

// Update crear nuevo usuario
func (self User) Update() (int64, error) {
	if self.Usuario == "" {
		return 0, errors.New("Usuario no especificado")
	}
	var users = make([]User, 0)
	if err := orm.Where("Usuario = ?", self.Usuario).Find(&users); err != nil {
		return 0, err
	}
	if _, err := govalidator.ValidateStruct(users[0]); err != nil {
		return 0, normalize(err, self)
	}
	users[0].Nombres = self.Nombres
	users[0].Apellidos = self.Apellidos

	affected, err := orm.Id(users[0].Id).Cols("nombres", "apellidos").Update(users[0])
	if err != nil {
		return affected, normalize(err, self)
	}
	return affected, nil
}
