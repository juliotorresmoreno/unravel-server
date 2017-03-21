package models

import "time"

import "errors"

import "../helper"
import "github.com/asaskevich/govalidator"

// User modelo de usuario
type User struct {
	Id        uint      `xorm:"bigint not null autoincr pk"`
	Nombres   string    `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Apellidos string    `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	FullName  string    `xorm:"varchar(200) not null" valid:"required,alphaSpaces"`
	Email     string    `xorm:"varchar(200) not null" valid:"required,email"`
	Usuario   string    `xorm:"varchar(100) not null unique index" valid:"required,username"`
	Passwd    string    `xorm:"varchar(100) not null" valid:"required,password"`
	Recovery  string    `xorm:"varchar(100) not null unique index"`
	Tipo      string    `xorm:"varchar(20) not null" valid:"required,alphanum"`
	Code      string    `xorm:"varchar(400) not null"`
	CreateAt  time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (that User) TableName() string {
	return "users"
}

func init() {
	orm.Sync2(new(User))
}

// Add crear nuevo usuario
func (that User) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(that)
	if err != nil {
		return 0, normalize(err, that)
	}
	that.Passwd = helper.Encript(that.Passwd)

	affected, err := orm.Insert(that)
	if err != nil {
		return affected, normalize(err, that)
	}
	return affected, nil
}

// ForceAdd crear nuevo usuario sin validar nada
func (that User) ForceAdd() (int64, error) {
	that.Passwd = helper.Encript(that.Passwd)
	affected, err := orm.Insert(that)
	if err != nil {
		return affected, normalize(err, that)
	}
	return affected, nil
}

// Update crear nuevo usuario
func (that User) Update() (int64, error) {
	if that.Usuario == "" {
		return 0, errors.New("Usuario no especificado")
	}
	var users = make([]User, 0)
	if err := orm.Where("Usuario = ?", that.Usuario).Find(&users); err != nil {
		return 0, err
	}
	users[0].Nombres = that.Nombres
	users[0].Apellidos = that.Apellidos
	if _, err := govalidator.ValidateStruct(users[0]); err != nil {
		return 0, normalize(err, that)
	}

	affected, err := orm.Id(users[0].Id).Cols("nombres", "apellidos").Update(users[0])
	if err != nil {
		return affected, normalize(err, that)
	}
	return affected, nil
}
