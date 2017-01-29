package social

import "time"

//Comentario estructura de un comentario
type Comentario struct {
	Usuario    string    `json:"usuario"`
	Comentario string    `json:"comentarios"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
