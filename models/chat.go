package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

// Chat modelo de los chats
type Chat struct {
	Id              uint      `xorm:"bigint not null autoincr pk"`
	UsuarioEmisor   string    `xorm:"varchar(100) not null index"`
	UsuarioReceptor string    `xorm:"varchar(100) not null index"`
	Message         string    `xorm:"text not null" valid:"required"`
	Leido           uint8     `json:"leido"`
	CreateAt        time.Time `xorm:"created"`
	UpdateAt        time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla del modelo
func (self Chat) TableName() string {
	return "chats"
}

func init() {
	orm.Sync2(new(Chat))
}

// Add agrega un nuevo chat
func (self Chat) Add() (int64, error) {
	if _, err := govalidator.ValidateStruct(self); err != nil {
		return 0, normalize(err, self)
	}
	var q = make([]User, 0)
	orm.Where("Usuario = ?", self.UsuarioReceptor).Find(&q)
	if len(q) == 1 {
		affected, err := orm.Insert(self)
		if err != nil {
			return affected, normalize(err, self)
		}
		return affected, nil
	}
	return 0, errors.New("El usuario no existe")
}
