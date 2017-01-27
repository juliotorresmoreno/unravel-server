package news

import "net/http"
import "../../models"
import "../../ws"

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = r.PostFormValue("usuario")
	var noticia = r.PostFormValue("noticia")
	var permiso = r.PostFormValue("permiso")

}
