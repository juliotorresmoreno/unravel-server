package news

import "net/http"
import "../../models"
import "../../ws"

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var noticia = r.PostFormValue("noticia")
	var permiso = r.PostFormValue("permiso")
	var lnoticia = models.Noticia{
		Usuario: session.Usuario,
		Noticia: noticia,
		Permiso: permiso,
	}
	if _, err := models.Add(lnoticia); err != nil {
		println(err.Error())
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{success:true}"))
}

// Listar listado de noticias en el muro
func Listar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{success:true}"))
}
