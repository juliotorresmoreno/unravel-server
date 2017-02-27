package social

import "time"

type Noticia struct {
	ID          interface{} "_id"
	Usuario     string      `valid:"required"`
	Nombres     string      `json:"nombres"`
	Apellidos   string      `json:"apellidos"`
	Noticia     string      `valid:"required"`
	Permiso     string      `valid:"required,matches(^(private|friends|public)$)"`
	Likes       []string
	Comentarios []Comentario
	CreateAt    time.Time
	UpdateAt    time.Time
}
