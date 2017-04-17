package news

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/social"
	"github.com/juliotorresmoreno/unravel-server/ws"
	"gopkg.in/mgo.v2/bson"
)

// Like el like de toda la vida
func Like(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	defer func() {
		if r := recover(); r != nil {
			helper.DespacharError(w, errors.New("Error desconocido"), http.StatusInternalServerError)
		}
	}()
	var ID = bson.ObjectIdHex(r.PostFormValue("noticia"))
	var friends, _ = models.GetFriends(session.Usuario)
	var length = len(friends)
	var usuarios = make([]string, length+1)
	for i := 0; i < length; i++ {
		usuarios[i] = friends[i].Usuario
	}
	usuarios[length] = session.Usuario
	var query = bson.M{
		"_id": ID,
		"usuario": map[string]interface{}{
			"$in": usuarios,
		},
	}
	var socialSS, SocialBD, err = social.GetSocial()
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	defer socialSS.Close()
	var resultado = noticia{}
	err = SocialBD.C(noticias).Find(query).One(&resultado)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	var existe = find(resultado.Likes, session.Usuario)
	if existe >= 0 {
		resultado.Likes = remove(resultado.Likes, existe)
	} else {
		resultado.Likes = append(resultado.Likes, session.Usuario)
	}
	var data = map[string]interface{}{
		"$set": map[string]interface{}{
			"likes":    resultado.Likes,
			"updateat": time.Now(),
		},
	}
	err = SocialBD.C(noticias).UpdateId(ID, data)
	if err != nil {
		helper.DespacharError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respuesta, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"data":    resultado,
	})
	w.Write(respuesta)
}
