package models

import "time"

// Profile modelo de usuario
type Profile struct {
	Id       uint      `xorm:"bigint not null pk"`
	Usuario  string    `xorm:"varchar(100) not null unique index" valid:"required"`
	Email    string    `xorm:"varchar(200)" valid:"email"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (u Profile) TableName() string {
	return "profile"
}

func init() {
	orm.Sync2(new(Profile))
}
