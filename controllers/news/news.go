package news

import "../../models"
import "encoding/json"
import "../../ws"
import "time"
import "net/http"
import "../../social"

// Publicar publica una noticia en el muro
func Publicar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var noticia = r.PostFormValue("noticia")
	var permiso = r.PostFormValue("permiso")
	var nueva = &social.Noticia{
		Usuario: session.Usuario,
		Noticia: noticia,
		Permiso: permiso,
	}
	w.Header().Set("Content-Type", "application/json")
	var err = social.Add("noticias", nueva)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		w.Write(respuesta)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"success\":true}"))
}

// PublicarOld publica una noticia en el muro
func PublicarOld(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
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
	w.Write([]byte("{\"success\":true}"))
}

// Listar listado de noticias en el muro
func Listar(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var _noticias = make([]models.Noticia, 0)
	var orm = models.GetXORM()
	var err = orm.Where("Usuario = ?", session.Usuario).OrderBy("create_at desc").Find(&_noticias)
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
	var length = len(_noticias)
	var _usuarios []string = make([]string, length)
	for i := 0; i < length; i++ {
		_usuarios[i] = _noticias[i].Usuario
	}
	usuarios, err := models.FindUsers(_usuarios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		w.Write(respuesta)
		return
	}
	var noticias = make([]noticia, length)
	for i := 0; i < length; i++ {
		for _, value := range usuarios {
			if _noticias[i].Usuario == value.Usuario {
				noticias[i] = noticia{
					Usuario:   value.Usuario,
					Nombres:   value.Nombres,
					Apellidos: value.Apellidos,
					Noticia:   _noticias[i].Noticia,
					Permiso:   _noticias[i].Permiso,
					CreateAt:  _noticias[i].CreateAt,
					UpdateAt:  _noticias[i].UpdateAt,
				}
			}
		}
	}
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    noticias,
	})
	w.Write(respuesta)
}

type noticia struct {
	Usuario   string    `json:"usuario"`
	Nombres   string    `json:"nombres"`
	Apellidos string    `json:"apellidos"`
	Noticia   string    `json:"noticia"`
	Permiso   string    `json:"permiso"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
