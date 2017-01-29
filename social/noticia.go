package social

import "time"

type Noticia struct {
	Usuario  string `valid:"required,alphanum"`
	Noticia  string `valid:"required,alphaSpaces"`
	Permiso  string `valid:"required,matches(^(private|friends|public)$)"`
	CreateAt time.Time
	UpdateAt time.Time
}
