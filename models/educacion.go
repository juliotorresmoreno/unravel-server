package models

import (
	"time"

	"github.com/juliotorresmoreno/unravel-server/db"
)

// Educacion modelo de los chats
type Educacion struct {
	ID uint `xorm:"id bigint not null autoincr pk" json:"id"`

	Usuario   string `xorm:"varchar(100) not null index" valid:"required" json:"usuario"`
	Pais      string `xorm:"varchar(200) not null" valid:"required" json:"pais"`
	Titulo    string `xorm:"varchar(100) not null" valid:"required" json:"titulo"`
	Grado     string `xorm:"varchar(100) not null" valid:"required" json:"grado" valid:"in(bachillerato,tecnico,licenciatura,profesional,especialidad,maestria,doctorado)"`
	AnoInicio string `xorm:"int not null" valid:"required" json:"ano_inicio"`
	AnoFin    string `xorm:"int not null" valid:"required" json:"ano_fin"`

	CreateAt time.Time `xorm:"created" json:"-"`
	UpdateAt time.Time `xorm:"updated" json:"-"`
}

//TableName establece el nombre de la tabla del modelo
func (el Educacion) TableName() string {
	return "educaciones"
}

func init() {
	var orm = db.GetXORM()
	orm.Sync2(new(Educacion))
	orm.Close()
}
