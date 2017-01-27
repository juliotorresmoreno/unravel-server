package models

import "time"

type Noticia struct {
	Usuario    string    `json:"usuario"`
	Nombres    string    `json:"nombres"`
	Apellidos  string    `json:"apellidos"`
	Estado     string    `json:"estado"`
	Registrado time.Time `json:"registrado"`
	Relacion   *Relacion `json:"relacion"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (that Noticia) TableName() string {
	return "noticias"
}
