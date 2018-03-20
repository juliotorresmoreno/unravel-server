package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/unravel-server/db"
)

// Chat modelo de los chats
type Chat struct {
	Id              uint      `xorm:"bigint not null autoincr pk"`
	Status          uint      `xorm:"int not null default 1"`
	UsuarioEmisor   string    `xorm:"varchar(100) not null index"`
	UsuarioReceptor string    `xorm:"varchar(100) not null index"`
	Message         string    `xorm:"text not null" valid:"required"`
	Leido           uint8     `json:"leido"`
	CreateAt        time.Time `xorm:"created"`
	UpdateAt        time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla del modelo
func (el Chat) TableName() string {
	return "chats"
}

func init() {
	var orm = db.GetXORM()
	orm.Sync2(new(Chat))
	orm.Close()
}

// Add agrega un nuevo chat
func (el Chat) Add() (uint, error) {
	if _, err := govalidator.ValidateStruct(el); err != nil {
		return 0, normalize(err, el)
	}
	var orm = db.GetXORM()
	defer orm.Close()
	length, _ := orm.Where("Usuario = ?", el.UsuarioReceptor).Count(&User{})
	if length == 1 {
		_, err := orm.Insert(&el)
		if err != nil {
			return 0, normalize(err, el)
		}
		return el.Id, nil
	}
	return 0, errors.New("El usuario no existe")
}

// Add agrega un nuevo chat
func (el Chat) Edit(id int) error {
	if _, err := govalidator.ValidateStruct(el); err != nil {
		return normalize(err, el)
	}
	var orm = db.GetXORM()
	defer orm.Close()
	length, _ := orm.Where("Usuario = ?", el.UsuarioReceptor).Count(&User{})
	if length == 1 {
		_, err := orm.Where("id = ?", id).Update(el)
		if err != nil {
			return normalize(err, el)
		}
		return nil
	}
	return errors.New("El usuario no existe")
}
