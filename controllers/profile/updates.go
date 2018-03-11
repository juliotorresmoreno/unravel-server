package profile

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func updateEmail(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_email") != "" && data.Get("email") != "" {
		if !helper.IsValidPermision(data.Get("permiso_email")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Email = data.Get("email")
		perfil.PermisoEmail = data.Get("permiso_email")
	}
	return nil
}

//
func updateNacimientoMes(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_nacimiento_dia") != "" && data.Get("nacimiento_dia") != "" &&
		data.Get("nacimiento_mes") != "" {
		if !helper.IsValidPermision(data.Get("permiso_nacimiento_dia")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.NacimientoDia = data.Get("nacimiento_dia")
		perfil.NacimientoMes = data.Get("nacimiento_mes")
		perfil.PermisoNacimientoDia = data.Get("permiso_nacimiento_dia")
	}
	return nil
}

func updateNacimientoAno(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_nacimiento_ano") != "" && data.Get("nacimiento_ano") != "" {
		if !helper.IsValidPermision(data.Get("permiso_nacimiento_ano")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.NacimientoAno = data.Get("nacimiento_ano")
		perfil.PermisoNacimientoAno = data.Get("permiso_nacimiento_ano")
	}
	return nil
}

func updateNacimientoSexo(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_sexo") != "" && data.Get("sexo") != "" {
		if !helper.IsValidPermision(data.Get("permiso_sexo")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Sexo = data.Get("sexo")
		perfil.PermisoSexo = data.Get("permiso_sexo")
	}
	return nil
}

func updateNacimientoPais(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_nacimiento_pais") != "" && data.Get("nacimiento_pais") != "" {
		if !helper.IsValidPermision(data.Get("permiso_nacimiento_pais")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.NacimientoPais = data.Get("nacimiento_pais")
		perfil.PermisoNacimientoPais = data.Get("permiso_nacimiento_pais")
	}
	return nil
}

func updateNacimientoCiudad(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_nacimiento_ciudad") != "" && data.Get("nacimiento_ciudad") != "" {
		if !helper.IsValidPermision(data.Get("permiso_nacimiento_ciudad")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.NacimientoCiudad = data.Get("nacimiento_ciudad")
		perfil.PermisoNacimientoCiudad = data.Get("permiso_nacimiento_ciudad")
	}
	return nil
}

func updateResidenciaPais(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_residencia_pais") != "" && data.Get("residencia_pais") != "" {
		if !helper.IsValidPermision(data.Get("permiso_residencia_pais")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.ResidenciaPais = data.Get("residencia_pais")
		perfil.PermisoResidenciaPais = data.Get("permiso_residencia_pais")
	}
	return nil
}

func updateResidenciaCiudad(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_residencia_ciudad") != "" && data.Get("residencia_ciudad") != "" {
		if !helper.IsValidPermision(data.Get("permiso_residencia_ciudad")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.ResidenciaCiudad = data.Get("residencia_ciudad")
		perfil.PermisoResidenciaCiudad = data.Get("permiso_residencia_ciudad")
	}
	return nil
}

func updateDireccion(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_direccion") != "" && data.Get("direccion") != "" {
		if !helper.IsValidPermision(data.Get("permiso_direccion")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Direccion = data.Get("direccion")
		perfil.PermisoDireccion = data.Get("permiso_direccion")
	}
	return nil
}

func updateTelefono(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_telefono") != "" && data.Get("telefono") != "" {
		if !helper.IsValidPermision(data.Get("permiso_telefono")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Telefono = data.Get("telefono")
		perfil.PermisoTelefono = data.Get("permiso_telefono")
	}
	return nil
}

func updateCelular(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_celular") != "" && data.Get("celular") != "" {
		if !helper.IsValidPermision(data.Get("permiso_celular")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Celular = data.Get("celular")
		perfil.PermisoCelular = data.Get("permiso_celular")
	}
	return nil
}

func updatePersonalidad(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_personalidad") != "" && data.Get("personalidad") != "" {
		if !helper.IsValidPermision(data.Get("permiso_personalidad")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Personalidad = data.Get("personalidad")
		perfil.PermisoPersonalidad = data.Get("permiso_personalidad")
	}
	return nil
}

func updateIntereses(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_intereses") != "" && data.Get("intereses") != "" {
		if !helper.IsValidPermision(data.Get("permiso_intereses")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Intereses = data.Get("intereses")
		perfil.PermisoIntereses = data.Get("permiso_intereses")
	}
	return nil
}

func updateMusica(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_musica") != "" && data.Get("musica") != "" {
		if !helper.IsValidPermision(data.Get("permiso_musica")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Musica = data.Get("musica")
		perfil.PermisoMusica = data.Get("permiso_musica")
	}
	return nil
}

func updateSeries(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_series") != "" && data.Get("series") != "" {
		if !helper.IsValidPermision(data.Get("permiso_series")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.Series = data.Get("series")
		perfil.PermisoSeries = data.Get("permiso_series")
	}
	return nil
}

func updateCreenciasReligiosas(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_creencias_religiosas") != "" && data.Get("creencias_religiosas") != "" {
		if !helper.IsValidPermision(data.Get("permiso_creencias_religiosas")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.CreenciasReligiosas = data.Get("creencias_religiosas")
		perfil.PermisoCreenciasReligiosas = data.Get("permiso_creencias_religiosas")
	}
	return nil
}

func updateCreenciasPoliticas(session *models.User, hub *ws.Hub, perfil *models.Profile, data url.Values) error {
	if data.Get("permiso_creencias_politicas") != "" && data.Get("creencias_politicas") != "" {
		if !helper.IsValidPermision(data.Get("permiso_creencias_politicas")) {
			return fmt.Errorf("El permiso especificado no es valido")
		}
		perfil.CreenciasPoliticas = data.Get("creencias_politicas")
		perfil.PermisoCreenciasPoliticas = data.Get("permiso_creencias_politicas")
	}
	return nil
}
