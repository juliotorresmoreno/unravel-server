package responses

import (
	"../../models"
	"time"
)

type Error struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Session struct {
	Usuario   string `json:"usuario"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Token     string `json:"token"`
}

type Login struct {
	Success bool    `json:"success"`
	Session Session `json:"session"`
}

type Success struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Mensaje struct {
	Action          string `json:"action"`
	Usuario   	string `json:"usuario"`
	UsuarioReceptor string `json:"usuarioReceptor"`
	Mensaje   	string `json:"mensaje"`
	Fecha     	time.Time  `json:"fecha"`
}

type SuccessData struct {
	Success bool      `json:"success"`
	Data    []Mensaje `json:"data"`
}

type ListFriends struct {
	Success bool `json:"success"`
	Data    []models.Friend `json:"data"`
}
