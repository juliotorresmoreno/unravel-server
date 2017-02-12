package category

import "net/http"
import "../../models"
import "../../ws"
import "../../helper"
import "encoding/json"

//GetCategorys Busqueda de personas
func GetCategorys(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var orm = models.GetXORM()
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
