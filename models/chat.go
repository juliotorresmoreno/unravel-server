package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Chat modelo de los chats
type Chat struct {
	Id              uint      `xorm:"bigint not null autoincr pk"`
	UsuarioEmisor   string    `xorm:"varchar(100) not null index"`
	UsuarioReceptor string    `xorm:"varchar(100) not null index"`
	Message         string    `xorm:"text not null" valid:"required"`
	Leido 		uint8     `json:"leido"`
	CreateAt        time.Time `xorm:"created"`
	UpdateAt        time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla del modelo
func (u Chat) TableName() string {
	return "chats"
}

func init() {
	orm.Sync2(new(Chat))
}

// Add agrega un nuevo chat
func (u Chat) Add() (int64, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return 0, normalize(err, u)
	}
	q := make([]User, 0)
	_ = orm.Where("Usuario = ?", u.UsuarioReceptor).Find(&q)
	if len(q) == 1 {
		affected, err := orm.Insert(u)
		if err != nil {
			return affected, normalize(err, u)
		}
		return affected, nil
	}
	return 0, werror{Msg: "El usuario no existe"}
}
