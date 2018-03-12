package profile

import (
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
)

type profile struct {
	models.Profile
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Estado    string `json:"estado"`
}

func truncar(p models.Profile, relacion int8) profile {
	var t = profile{}

	t.Estado = map[int8]string{
		models.EstadoAceptado:    "Amigos",
		models.EstadoSolicitado:  "Solicitado",
		models.EstadoDesconocido: "Desconocido",
	}[relacion]

	t.Usuario = p.Usuario
	t.Legenda = p.Legenda
	t.Descripcion = p.Descripcion
	t.PrecioHora = p.PrecioHora

	if helper.PuedoVer(relacion, p.PermisoEmail) {
		t.Email = p.Email
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoDia) {
		t.NacimientoMes = p.NacimientoMes
		t.NacimientoDia = p.NacimientoDia
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoAno) {
		t.NacimientoAno = p.NacimientoAno
	}
	if helper.PuedoVer(relacion, p.PermisoSexo) {
		t.Sexo = p.Sexo
	}

	if helper.PuedoVer(relacion, p.PermisoNacimientoPais) {
		t.NacimientoPais = p.NacimientoPais
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoCiudad) {
		t.NacimientoCiudad = p.NacimientoCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaPais) {
		t.ResidenciaPais = p.ResidenciaPais
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaCiudad) {
		t.ResidenciaCiudad = p.ResidenciaCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoDireccion) {
		t.Direccion = p.Direccion
	}
	if helper.PuedoVer(relacion, p.PermisoTelefono) {
		t.Telefono = p.Telefono
	}
	if helper.PuedoVer(relacion, p.PermisoCelular) {
		t.Celular = p.Celular
	}

	if helper.PuedoVer(relacion, p.PermisoPersonalidad) {
		t.Personalidad = p.Personalidad
	}
	if helper.PuedoVer(relacion, p.PermisoIntereses) {
		t.Intereses = p.Intereses
	}
	if helper.PuedoVer(relacion, p.PermisoSeries) {
		t.Series = p.Series
	}
	if helper.PuedoVer(relacion, p.PermisoMusica) {
		t.Musica = p.Musica
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasReligiosas) {
		t.CreenciasReligiosas = p.CreenciasReligiosas
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasPoliticas) {
		t.CreenciasPoliticas = p.CreenciasPoliticas
	}

	return t
}
