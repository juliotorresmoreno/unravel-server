package responses

import "../../models"

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
	Usuario   string `json:"usuario"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Mensaje   string `json:"mensaje"`
	Fecha     int64  `json:"fecha"`
}

type SuccessData struct {
	Success bool      `json:"success"`
	Data    []Mensaje `json:"data"`
}

type ListFriends struct {
	Success bool `json:"success"`
	Data    []models.Friend `json:"data"`
}
