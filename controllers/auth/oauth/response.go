package oauth

//Usuario resultado de la autenticacion oauth
type Usuario struct {
	Nombres   string
	Apellidos string
	FullName  string
	Email     string
	Tipo      string
	Usuario   string
	Code      string
}
