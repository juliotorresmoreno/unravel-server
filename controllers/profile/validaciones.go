package profile

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

type responseError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type responseSuccess struct {
	Success bool `json:"success"`
}

// Update actualiza los datos del perfil
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	perfil := getProfile(session)
	data := helper.GetPostParams(r)
	if data.Get("nombres") != "" && data.Get("apellidos") != "" {
		updateProfile(w, r, session, hub, data)
		return
	}

	if err := updateEmail(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateNacimientoMes(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateNacimientoAno(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateNacimientoSexo(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateNacimientoPais(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateNacimientoCiudad(w, r, session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateResidenciaPais(w, r, session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateResidenciaCiudad(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateDireccion(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateTelefono(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateCelular(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updatePersonalidad(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateIntereses(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateSeries(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateMusica(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateCreenciasReligiosas(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateCreenciasPoliticas(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	if err := updateAll(session, hub, perfil, data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
