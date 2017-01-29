package social

import "time"

type Chat struct {
	UsuarioEmisor   string
	UsuarioReceptor string
	Message         string
	Leido           uint8
	CreateAt        time.Time
	UpdateAt        time.Time
}
