package models

import (
	"github.com/asaskevich/govalidator"
	"../services"
)

type User struct {
	Id		uint `xorm:"bigint not null primary"`
	Nombres 	string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Apellidos	string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Email 		string `xorm:"varchar(200) not null unique" gorm:"unique" valid:"required,email"`
	Usuario		string `xorm:"varchar(100) not null unique" valid:"required,alphanum"`
	Passwd		string `xorm:"varchar(100) not null" valid:"required,password"`
}

func init() {
	orm.Sync2(new(User))
}

func(u User) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return 0, normalize(err, u)
	}
	u.Passwd = services.Encript(u.Passwd)

	affected, err := orm.Insert(u)
	if err != nil {
		return affected, normalize(err, u)
	} else {
		return affected, nil
	}
}

