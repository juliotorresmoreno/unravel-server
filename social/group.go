package social

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Group struct {
	ID          bson.ObjectId "_id"
	Usuario     string        `json:"usuario"`
	Nombre      string        `json:"nombre" valid:"required,alphaSpaces"`
	Descripcion string        `json:"descripcion"`
	Categoria   int           `json:"categoria"`
	Permiso     string        `json:"permiso" valid:"matches(^(private|friends|public)$)`
	CreateAt    time.Time
	UpdateAt    time.Time
}
