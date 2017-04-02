package models

import "time"

type Noticia struct {
	Id       uint      `xorm:"bigint not null autoincr pk" json:"-"`
	Usuario  string    `xorm:"varchar(100) not null index" valid:"required" json:"usuario"`
	Noticia  string    `xorm:"text not null" valid:"required,password" json:"noticia"`
	Permiso  string    `xorm:"varchar(10) not null" valid:"required,matches(^(private|friends|public)$) json:"permiso"`
	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (that Noticia) TableName() string {
	return "noticias"
}

func init() {
	var orm = GetXORM()
	orm.Sync2(new(Noticia))
	orm.Close()
}
