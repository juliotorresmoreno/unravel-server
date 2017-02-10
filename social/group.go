package social

import "time"

type Group struct {
	Usuario     string `json:"usuario"`
	Nombre      string `json:"nombre" valid:"required,alphaSpaces"`
	Descripcion string `json:"descripcion"`
	Permiso     string `json:"permiso" valid:"matches(^(private|friends|public)$)`
	CreateAt    time.Time
	UpdateAt    time.Time
}
