package models

import (
	"time"

	"github.com/juliotorresmoreno/unravel-server/db"
)

// Skill modelo de los chats
type Skill struct {
	ID uint `xorm:"id bigint not null autoincr pk" json:"id"`

	Usuario string `xorm:"varchar(100) not null index" valid:"required" json:"usuario"`
	Nombre  string `xorm:"varchar(200) not null" valid:"required" json:"nombre"`

	CreateAt time.Time `xorm:"created" json:"-"`
	UpdateAt time.Time `xorm:"updated" json:"-"`
}

//TableName establece el nombre de la tabla del modelo
func (el Skill) TableName() string {
	return "skills"
}

func init() {
	var orm = db.GetXORM()
	orm.Sync2(new(Skill))
	orm.Close()
}
