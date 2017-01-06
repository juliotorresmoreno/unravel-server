package models

import "time"

const (
	EstadoSolicitado = 0
	EstadoAceptado   = 1
)

// Relacion modelo de usuario
type Relacion struct {
	Id                uint      `xorm:"bigint not null autoincr pk" json:"id"`
	UsuarioSolicita   string    `xorm:"varchar(100) not null index" valid:"required" json:"usuario_solicita"`
	UsuarioSolicitado string    `xorm:"varchar(100) not null index" valid:"required" json:"usuario_solicitado"`
	EstadoRelacion    uint8     `xorm:"tinyint not null"`
	CreateAt          time.Time `xorm:"created" json:"create_at"`
	UpdateAt          time.Time `xorm:"updated" json:"update_at"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (u Relacion) TableName() string {
	return "relacion"
}

func init() {
	orm.Sync2(new(Relacion))
}
