package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
)

// User modelo de usuario
type User struct {
	Id        uint   `xorm:"bigint not null autoincr pk"`
	Nombres   string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Apellidos string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	FullName  string `xorm:"varchar(200) not null" valid:"required,alphaSpaces"`
	Email     string `xorm:"varchar(200) not null" valid:"required,email"`
	Usuario   string `xorm:"varchar(100) not null unique index" valid:"required,username"`
	Passwd    string `xorm:"varchar(100) not null" valid:"required,password,encript"`
	Recovery  string `xorm:"varchar(100) not null index"`
	Tipo      string `xorm:"varchar(20) not null" valid:"required,alphanum"`
	Code      string `xorm:"varchar(400) not null"`

	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (el User) TableName() string {
	return "users"
}

func init() {
	var orm = db.GetXORM()
	orm.Sync2(new(User))
	orm.Close()
}

// Add crear nuevo usuario
func (el User) Add() (int64, error) {
	_, err := helper.ValidateStruct(el)
	if err != nil {
		return 0, normalize(err, el)
	}
	var orm = db.GetXORM()
	defer orm.Close()
	el.Passwd = helper.Encript(el.Passwd)

	affected, err := orm.Insert(el)
	if err != nil {
		return affected, normalize(err, el)
	}
	return affected, nil
}

// ForceAdd crear nuevo usuario sin validar nada
func (el User) ForceAdd() (int64, error) {
	if el.Passwd != "" {
		el.Passwd = helper.Encript(el.Passwd)
	}
	var orm = db.GetXORM()
	defer orm.Close()
	affected, err := orm.Insert(el)
	if err != nil {
		return affected, normalize(err, el)
	}
	return affected, nil
}

// Update crear nuevo usuario
func (el User) Update() (int64, error) {
	if el.Usuario == "" {
		return 0, errors.New("Usuario no especificado")
	}
	var orm = db.GetXORM()
	defer orm.Close()
	var users = make([]User, 0)
	if err := orm.Where("Usuario = ?", el.Usuario).Find(&users); err != nil {
		return 0, err
	}
	users[0].Nombres = el.Nombres
	users[0].Apellidos = el.Apellidos
	users[0].FullName = el.FullName
	if _, err := govalidator.ValidateStruct(users[0]); err != nil {
		return 0, normalize(err, el)
	}
	affected, err := orm.Id(users[0].Id).Cols("nombres", "apellidos", "full_name").Update(users[0])
	if err != nil {
		return affected, normalize(err, el)
	}
	return affected, nil
}
