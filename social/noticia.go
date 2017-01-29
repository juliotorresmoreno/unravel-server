package social

import "time"

type Noticia struct {
	Usuario     string `valid:"required,alphanum"`
	Noticia     string `valid:"required"`
	Permiso     string `valid:"required,matches(^(private|friends|public)$)"`
	Likes       []string
	Comentarios []string
	CreateAt    time.Time
	UpdateAt    time.Time
}
