package helper

import (
	"encoding/json"
	"net/http"
)

//DespacharError responde los errores
func DespacharError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	})
	w.Write(respuesta)
}
