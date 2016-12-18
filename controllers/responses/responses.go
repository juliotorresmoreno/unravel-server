package responses

type Error struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type Session struct {
	Usuario string `json:"usuario"`
	Nombres string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Token string `json:"token"`
}

type Login struct {
	Success bool `json:"success"`
	Session Session `json:"session"`
}