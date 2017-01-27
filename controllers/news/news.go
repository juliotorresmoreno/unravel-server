package news

import "net/http"
import "../../models"
import "../../ws"
import "encoding/json"

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var noticia = r.PostFormValue("noticia")
	var permiso = r.PostFormValue("permiso")
	var lnoticia = models.Noticia{
		Usuario: session.Usuario,
		Noticia: noticia,
		Permiso: permiso,
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := models.Add(lnoticia); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		w.Write(respuesta)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{success:true}"))
}

// Listar listado de noticias en el muro
func Listar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var noticia = make([]models.Noticia, 0)
	var orm = models.GetXORM()
	var err = orm.Where("Usuario = ?", session.Usuario).Find(&noticia)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		w.Write(respuesta)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    noticia,
	})
	w.Write(respuesta)
}
