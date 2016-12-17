package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

//xorm:"lusers"
type Users struct {
	gorm.Model
	Id uint `gorm:"primary_key"`
	Nombres string `valid:"required,alphaSpaces"`
	Apellidos string `valid:"required,alphaSpaces"`
	Email string `gorm:"unique" valid:"required,email"`
	Usuario string `gorm:"unique" valid:"required,alphanum"`
	Passwd string `valid:"required,password"`
}

func init() {
	engine.AutoMigrate(&Users{})
	engine.Model(&Users{}).Update("CreatedAt", time.Now())
}

func(u Users) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return 0, err
	}
	password := []byte(u.Passwd)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	u.Passwd = string(hashedPassword)

	resp:= engine.Create(&u)
	if(len(resp.GetErrors()) > 0) {
		return resp.RowsAffected, normalize(resp.GetErrors()[0], u)
	} else {
		return resp.RowsAffected, nil
	}
}