package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func protect(fn func(w http.ResponseWriter, r *http.Request, user *models.User, hub *ws.Hub), hub *ws.Hub, rechazar bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cache = db.GetCache()
		var token = helper.GetToken(r)
		var session = cache.Get(token)
		var usuario, _ = session.Result()
		var users = make([]models.User, 0)
		var orm = db.GetXORM()
		defer orm.Close()

		helper.Cors(w, r)

		if session.Err() != nil {
			if rechazar {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				fn(w, r, nil, hub)
			}
			return
		}

		if err := orm.Where("Usuario = ?", usuario).Find(&users); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			respuesta, _ := json.Marshal(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			w.Write(respuesta)
			return
		}
		if len(users) == 1 {
			cache := db.GetCache()
			cache.Set(
				string(token),
				users[0].Usuario,
				time.Duration(config.SESSION_DURATION)*time.Second,
			)
			http.SetCookie(w, &http.Cookie{
				HttpOnly: true,
				MaxAge:   config.SESSION_DURATION,
				Name:     "token",
				Value:    token,
				Path:     "/",
			})
			fn(w, r, &users[0], hub)
		} else {
			if rechazar {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				fn(w, r, nil, hub)
			}
		}
	}
}
