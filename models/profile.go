package models

import "time"

// Profile modelo de usuario
type Profile struct {
	Id                   uint   `xorm:"bigint not null autoincr pk" json:"id"`
	Usuario              string `xorm:"varchar(100) not null unique index" valid:"required" json:"usuario"`
	Email                string `xorm:"varchar(200)" valid:"email" json:"email"`
	PermisoEmail         string `xorm:"varchar(20)" json:"permiso_email"`
	NacimientoDia        string `xorm:"varchar(2)" json:"nacimiento_dia"`
	NacimientoMes        string `xorm:"varchar(2)" json:"nacimiento_mes"`
	PermisoNacimientoDia string `xorm:"varchar(20)" json:"permiso_email"`

	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (u Profile) TableName() string {
	return "profile"
}

func init() {
	orm.Sync2(new(Profile))
}
