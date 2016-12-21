package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"../controllers/auth"
	"../controllers/friends"
	"../controllers/chats"
	"../controllers/responses"
	"../test"
	"../models"
	"../helper"
	"../config"
	"../ws"
	"encoding/json"
	"time"
)

func protect(fn func(w http.ResponseWriter, r *http.Request, user *models.User, hub *ws.Hub), hub *ws.Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cache := models.GetCache()
		_token := helper.GetCookie(r, "token")
		if _token == "" {
			_token = r.URL.Query().Get("token")
		}
		session := cache.Get(_token)
		usuario, _ := session.Result()

		if session.Err() != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		users := make([]models.User, 0)
		orm := models.GetXORM()

		err := orm.Where("Usuario = ?", usuario).Find(&users)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			respuesta, _ := json.Marshal(responses.Error{Success:false,Error:err.Error()})
			w.Write(respuesta)
			return
		}
		if len(users) == 1 {
			cache := models.GetCache()
			cache.Set(
				string(_token),
				users[0].Usuario,
				time.Duration(config.SESSION_DURATION) * time.Second,
			)
			http.SetCookie(w, &http.Cookie{
				HttpOnly: true,
				MaxAge: config.SESSION_DURATION,
				Name: "token",
				Value: _token,
				Path: "/",
			})
			fn(w, r, &users[0], hub)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	}
}

func GetHandler() http.Handler {
	webroot := http.FileServer(http.Dir("webroot"))
	mux := mux.NewRouter().StrictSlash(false)
	hub := ws.GetHub()

	// auth
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")
	mux.HandleFunc("/api/v1/auth/login", auth.Login).Methods("POST")
	mux.HandleFunc("/api/v1/auth/session", auth.Session).Methods("GET")

	mux.HandleFunc("/api/v1/friends", protect(friends.ListFriends, hub)).Methods("GET")
	mux.HandleFunc("/api/v1/chats/{user}", protect(chats.List, hub)).Methods("GET")
	mux.HandleFunc("/api/v1/chats/mensaje", protect(chats.Mensaje, hub)).Methods("POST")

	mux.HandleFunc("/test", test.Test).Methods("GET")


	mux.HandleFunc("/ws", protect(func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		ws.ServeWs(hub, w, r, session)
	}, hub))

	mux.Handle("/", webroot)
	mux.Handle("/{all}", webroot)
	return mux
}
