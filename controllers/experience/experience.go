package experience

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/middlewares"

	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

type responseSuccess struct {
	Success bool `json:"success"`
}

type responseError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type responseData struct {
	Success bool                `json:"success"`
	Data    []models.Experience `json:"data"`
}

//NewRouter hola mundo
func NewRouter(hub *ws.Hub) http.Handler {
	//session := &models.User{}
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", middlewares.Protect(Read, hub, true)).Methods("GET")
	router.HandleFunc("/", middlewares.Protect(Create, hub, true)).Methods("POST")
	router.HandleFunc("/{id}", middlewares.Protect(Update, hub, true)).Methods("PUT")
	router.HandleFunc("/{id}", middlewares.Protect(Delete, hub, true)).Methods("DELETE")
	return router
}

// Read una nueva experiencia laboral
func Read(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	experiences := make([]models.Experience, 0)
	orm := db.GetXORM()
	defer orm.Close()

	err := orm.Where("usuario = ?", session.Usuario).Find(&experiences)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData{
		Success: true,
		Data:    experiences,
	})
}

// Create una nueva experiencia laboral
func Create(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)

	experience := models.Experience{}
	experience.Usuario = session.Usuario
	experience.Cargo = data.Get("cargo")
	experience.Empresa = data.Get("empresa")
	experience.AnoInicio = data.Get("ano_inicio")
	experience.MesInicio = data.Get("mes_inicio")
	experience.ContinuoTrabajando = data.Get("continuo_trabajando")
	experience.AnoFin = data.Get("ano_inicio")
	experience.MesFin = data.Get("mes_fin")

	w.Header().Set("Content-Type", "application/json")
	if _, err := models.Add(experience); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuccess{
		Success: true,
	})
}

// Update una nueva experiencia laboral
func Update(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)

	vars := mux.Vars(r)

	_id, _ := strconv.Atoi(vars["id"])
	id := uint(_id)

	experience := models.Experience{}
	experience.Usuario = session.Usuario
	experience.Cargo = data.Get("cargo")
	experience.Empresa = data.Get("empresa")
	experience.AnoInicio = data.Get("ano_inicio")
	experience.MesInicio = data.Get("mes_inicio")
	experience.ContinuoTrabajando = data.Get("continuo_trabajando")
	experience.AnoFin = data.Get("ano_fin")
	experience.MesFin = data.Get("mes_fin")

	w.Header().Set("Content-Type", "application/json")
	if _, err := models.Update(id, experience); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuccess{
		Success: true,
	})
}

// Delete una nueva experiencia laboral
func Delete(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	orm := db.GetXORM()
	defer orm.Close()

	w.Header().Set("Content-Type", "application/json")
	if _, err := orm.Delete(models.Experience{ID: 0}); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuccess{
		Success: true,
	})
}
