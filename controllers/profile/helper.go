package profile

import "../../models"
import "../../helper"

func truncar(p models.Profile, relacion int8) map[string]string {
	var r = make(map[string]string)

	r["estado"] = map[int8]string{
		models.EstadoAceptado:    "Amigos",
		models.EstadoSolicitado:  "Solicitado",
		models.EstadoDesconocido: "Desconocido",
	}[relacion]

	if helper.PuedoVer(relacion, p.PermisoEmail) {
		r["email"] = p.Email
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoDia) {
		r["nacimiento_mes"] = p.NacimientoMes
		r["nacimiento_dia"] = p.NacimientoDia
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoAno) {
		r["nacimiento_ano"] = p.NacimientoAno
	}
	if helper.PuedoVer(relacion, p.PermisoSexo) {
		r["sexo"] = p.Sexo
	}

	if helper.PuedoVer(relacion, p.PermisoNacimientoPais) {
		r["nacimiento_pais"] = p.NacimientoPais
	}
	if helper.PuedoVer(relacion, p.PermisoNacimientoCiudad) {
		r["nacimiento_ciudad"] = p.NacimientoCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaPais) {
		r["residencia_pais"] = p.ResidenciaPais
	}
	if helper.PuedoVer(relacion, p.PermisoResidenciaCiudad) {
		r["residencia_ciudad"] = p.ResidenciaCiudad
	}
	if helper.PuedoVer(relacion, p.PermisoDireccion) {
		r["direccion"] = p.Direccion
	}
	if helper.PuedoVer(relacion, p.PermisoTelefono) {
		r["telefono"] = p.Telefono
	}
	if helper.PuedoVer(relacion, p.PermisoCelular) {
		r["celular"] = p.Celular
	}

	if helper.PuedoVer(relacion, p.PermisoPersonalidad) {
		r["personalidad"] = p.Personalidad
	}
	if helper.PuedoVer(relacion, p.PermisoIntereses) {
		r["intereses"] = p.Intereses
	}
	if helper.PuedoVer(relacion, p.PermisoSeries) {
		r["series"] = p.Series
	}
	if helper.PuedoVer(relacion, p.PermisoMusica) {
		r["musica"] = p.Musica
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasReligiosas) {
		r["creencias_religiosas"] = p.CreenciasReligiosas
	}
	if helper.PuedoVer(relacion, p.PermisoCreenciasPoliticas) {
		r["creencias_politicas"] = p.CreenciasPoliticas
	}

	return r
}
