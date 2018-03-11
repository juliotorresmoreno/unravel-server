package models

import (
	"time"

	"github.com/juliotorresmoreno/unravel-server/db"
)

// Experience modelo de los chats
type Experience struct {
	ID uint `xorm:"id bigint not null autoincr pk" json:"id"`

	Usuario            string `xorm:"varchar(100) not null index" valid:"required" json:"usuario"`
	Cargo              string `xorm:"varchar(200) not null" valid:"required" json:"cargo"`
	Empresa            string `xorm:"varchar(100) not null" valid:"required" json:"empresa"`
	AnoInicio          string `xorm:"int not null" valid:"required" json:"ano_inicio"`
	MesInicio          string `xorm:"int not null" valid:"required" json:"mes_inicio"`
	ContinuoTrabajando string `xorm:"int not null" valid:"required" json:"continuo_trabajando"`
	AnoFin             string `xorm:"int not null" valid:"required" json:"ano_fin"`
	MesFin             string `xorm:"int not null" valid:"required" json:"mes_fin"`

	CreateAt time.Time `xorm:"created" json:"-"`
	UpdateAt time.Time `xorm:"updated" json:"-"`
}

//TableName establece el nombre de la tabla del modelo
func (el Experience) TableName() string {
	return "experience"
}

func init() {
	var orm = db.GetXORM()
	orm.Sync2(new(Experience))
	orm.Close()
}
