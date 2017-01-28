package social

import "time"

type Noticia struct {
	Id       uint
	Usuario  string
	Noticia  string
	Permiso  string
	CreateAt time.Time
	UpdateAt time.Time
}

