package models

import "time"

type Noticia struct {
	Id       uint      `xorm:"bigint not null autoincr pk"`
	Usuario  string    `xorm:"varchar(100) not null index" valid:"required,alphanum"`
	Noticia  string    `xorm:"text not null" valid:"required,password"`
	Permiso  string    `xorm:"varchar(10) not null" valid:"required"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (that Noticia) TableName() string {
	return "noticias"
}

func init() {
	orm.Sync2(new(Noticia))
}
