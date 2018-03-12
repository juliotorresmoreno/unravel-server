package skill

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
	Success bool           `json:"success"`
	Data    []models.Skill `json:"data"`
}

//NewRouter hola mundo
func NewRouter(hub *ws.Hub) http.Handler {
	//session := &models.User{}
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", middlewares.Protect(Read, hub, true)).Methods("GET")
	router.HandleFunc("/{username}", middlewares.Protect(Read, hub, true)).Methods("GET")
	router.HandleFunc("/", middlewares.Protect(Create, hub, true)).Methods("POST")
	router.HandleFunc("/{id}", middlewares.Protect(Update, hub, true)).Methods("PUT")
	router.HandleFunc("/{id}", middlewares.Protect(Delete, hub, true)).Methods("DELETE")
	return router
}

// Read una nueva experiencia laboral
func Read(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var vars = mux.Vars(r)
	var usuario string
	if vars["username"] != "" {
		usuario = vars["username"]
	} else {
		usuario = session.Usuario
	}

	educacions := make([]models.Skill, 0)
	orm := db.GetXORM()
	defer orm.Close()

	err := orm.Where("usuario = ?", usuario).Find(&educacions)
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
		Data:    educacions,
	})
}

// Create una nueva experiencia laboral
func Create(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	data := helper.GetPostParams(r)

	skill := models.Skill{}
	skill.Usuario = session.Usuario
	skill.Nombre = data.Get("nombre")

	w.Header().Set("Content-Type", "application/json")
	if _, err := models.Add(skill); err != nil {
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

	skill := models.Skill{}
	skill.Usuario = session.Usuario
	skill.Nombre = data.Get("nombre")

	orm := db.GetXORM()
	defer orm.Close()

	usuario := session.Usuario
	if c, _ := orm.ID(id).Where("usuario = ?", usuario).Count(models.Skill{}); c == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responseError{Error: "Acceso no autorizado"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := models.Update(id, skill); err != nil {
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

	vars := mux.Vars(r)

	_id, _ := strconv.Atoi(vars["id"])
	id := uint(_id)

	usuario := session.Usuario
	if c, _ := orm.ID(id).Where("usuario = ?", usuario).Count(models.Skill{}); c == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responseError{Error: "Acceso no autorizado"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := orm.ID(id).Delete(models.Skill{ID: id}); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuccess{
		Success: true,
	})
}
