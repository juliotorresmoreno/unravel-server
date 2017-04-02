package models

import (
	"time"
)

// Category modelo de los chats
type Category struct {
	Id       uint      `xorm:"bigint not null autoincr pk" json:"id"`
	Nombre   string    `xorm:"varchar(200) not null" valid:"required" json:"nombre"`
	Lang     string    `xorm:"varchar(5) not null" valid:"required" json:"lang"`
	CreateAt time.Time `xorm:"created" json:"-"`
	UpdateAt time.Time `xorm:"updated" json:"-"`
}

//TableName establece el nombre de la tabla del modelo
func (el Category) TableName() string {
	return "categorys"
}

func init() {
	var orm = GetXORM()
	orm.Sync2(new(Category))
	orm.Close()
}
