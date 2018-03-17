package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/middlewares"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

type responseError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type responseData struct {
	Success bool                `json:"success"`
	Data    []models.Experience `json:"data"`
}

type ciudades map[string][]string

//NewRouter hola mundo
func NewRouter(hub *ws.Hub) http.Handler {
	data := "./data/geo.json"
	content, err := ioutil.ReadFile(data)
	if err != nil {
		fmt.Println(err)
	}
	ciudades := ciudades{}
	if err = json.Unmarshal(content, &ciudades); err != nil {
		fmt.Println(err)
	}
	paises := make([]string, 0)
	for name := range ciudades {
		paises = append(paises, name)
	}
	Read := NewReader(paises, ciudades)
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", middlewares.Protect(Read, hub, true)).Methods("GET")
	router.HandleFunc("/{country}", middlewares.Protect(Read, hub, true)).Methods("GET")
	return router
}

//HandlerFunc d
type HandlerFunc func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub)

// NewReader una nueva experiencia laboral
func NewReader(paises []string, ciudades ciudades) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		var vars = mux.Vars(r)
		country := vars["country"]
		w.Header().Set("Content-Type", "application/json")
		if country == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"data":    paises,
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data":    ciudades[country],
		})
	}
}
