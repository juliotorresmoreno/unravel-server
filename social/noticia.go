package social

import "time"

type Noticia struct {
	Usuario     string `valid:"required,alphanum"`
	Nombres     string `json:"nombres"`
	Apellidos   string `json:"apellidos"`
	Noticia     string `valid:"required"`
	Permiso     string `valid:"required,matches(^(private|friends|public)$)"`
	Likes       []string
	Comentarios []string
	CreateAt    time.Time
	UpdateAt    time.Time
}
