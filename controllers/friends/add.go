package friends

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func obtenerRelaciones(session, usuario string) ([]models.Relacion, error) {
	var relaciones = make([]models.Relacion, 0)
	var orm = db.GetXORM()
	defer orm.Close()
	var str = "(usuario_solicita = ? and usuario_solicitado = ?) or (usuario_solicita = ? and usuario_solicitado = ?)"
	err := orm.Where(str, usuario, session, session, usuario).Find(&relaciones)
	return relaciones, err
}

// Add agregar amigo
func Add(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	var usuario = r.PostFormValue("user")
	if usuario == "" || usuario == session.Usuario {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var relacion *models.Relacion
	relaciones, err := obtenerRelaciones(session.Usuario, usuario)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Existe la solicitud de amistad o la amistad en si
	if len(relaciones) == 1 {
		// Es una peticion hecha por la otra parte
		if relaciones[0].EstadoRelacion == 0 && relaciones[0].UsuarioSolicita == usuario {
			relaciones[0].EstadoRelacion = 1
			models.Update(relaciones[0].Id, relaciones[0])
		}
		relacion = &relaciones[0]
	} else {
		relacion = &models.Relacion{
			UsuarioSolicita:   session.Usuario,
			UsuarioSolicitado: usuario,
			EstadoRelacion:    0,
		}
		models.Add(relacion)
	}
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success":  true,
		"relacion": relacion,
		"estado":   "Solicitado",
	})
	w.Write(respuesta)
}
