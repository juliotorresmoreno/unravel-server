package router

import (
	"encoding/json"
	"net/http"
	"time"

	"../config"
	"../helper"
	"../models"
	"../ws"
)

func protect(fn func(w http.ResponseWriter, r *http.Request, user *models.User, hub *ws.Hub), hub *ws.Hub, rechazar bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cache = models.GetCache()
		var _token = helper.GetCookie(r, "token")
		if _token == "" {
			_token = r.URL.Query().Get("token")
		}
		var session = cache.Get(_token)
		var usuario, _ = session.Result()

		if session.Err() != nil {
			if rechazar {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				fn(w, r, nil, hub)
			}
			return
		}

		var users = make([]models.User, 0)
		var orm = models.GetXORM()

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
			cache := models.GetCache()
			cache.Set(
				string(_token),
				users[0].Usuario,
				time.Duration(config.SESSION_DURATION)*time.Second,
			)
			http.SetCookie(w, &http.Cookie{
				HttpOnly: true,
				MaxAge:   config.SESSION_DURATION,
				Name:     "token",
				Value:    _token,
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
