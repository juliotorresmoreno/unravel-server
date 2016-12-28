package models

import (
	"time"

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
	orm.Sync2(new(Chat))
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
