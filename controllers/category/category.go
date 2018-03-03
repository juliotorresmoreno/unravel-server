package category

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

//GetCategorys Busqueda de personas
func GetCategorys(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var orm = db.GetXORM()
	defer orm.Close()
	var resultado = make([]models.Category, 0)
	var err = orm.Find(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	var respuesta, _ = json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)
}
