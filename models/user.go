package models

import (
	"github.com/asaskevich/govalidator"
	"../helper"
)

type User struct {
	Id		uint `xorm:"bigint not null autoincr pk"`
	Nombres 	string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Apellidos	string `xorm:"varchar(100) not null" valid:"required,alphaSpaces"`
	Email 		string `xorm:"varchar(200) not null unique" gorm:"unique" valid:"required,email"`
	Usuario		string `xorm:"varchar(100) not null unique index" valid:"required,alphanum"`
	Passwd		string `xorm:"varchar(100) not null" valid:"required,password"`
}

func(u User) TableName() string {
	return "users"
}

func init() {
	orm.Sync2(new(User))
}

func(u User) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return 0, normalize(err, u)
	}
	u.Passwd = helper.Encript(u.Passwd)

	affected, err := orm.Insert(u)
	if err != nil {
		return affected, normalize(err, u)
	} else {
		return affected, nil
	}
}

