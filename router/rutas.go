package router

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"../config"
	"../controllers/auth"
	"../controllers/chats"
	"../controllers/friends"
	"../controllers/profile"
	"../controllers/responses"
	"../helper"
	"../models"
	"../test"
	"../ws"
	"github.com/gorilla/mux"
)

func protect(fn func(w http.ResponseWriter, r *http.Request, user *models.User, hub *ws.Hub), hub *ws.Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cache = models.GetCache()
		var _token = helper.GetCookie(r, "token")
		if _token == "" {
			_token = r.URL.Query().Get("token")
		}
		var session = cache.Get(_token)
		var usuario, _ = session.Result()

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
			respuesta, _ := json.Marshal(responses.Error{Success: false, Error: err.Error()})
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
			w.WriteHeader(http.StatusUnauthorized)
		}

	}
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	var publicPath = "./webroot"
	var path = publicPath + "/" + r.URL.Path
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.ServeFile(w, r, publicPath+"/index.html")
}

// GetHandler aca se establecen las rutas del router
func GetHandler() http.Handler {
	var mux = mux.NewRouter().StrictSlash(false)
	var hub = ws.GetHub()

	// auth
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")
	mux.HandleFunc("/api/v1/auth/login", auth.Login).Methods("POST")
	mux.HandleFunc("/api/v1/auth/session", auth.Session).Methods("GET")
	mux.HandleFunc("/api/v1/auth/logout", auth.Logout).Methods("GET")

	// profile
	mux.HandleFunc("/api/v1/profile", protect(profile.Update, hub)).Methods("POST", "PUT")

	// friends
	mux.HandleFunc("/api/v1/friends", protect(friends.ListFriends, hub)).Methods("GET")

	// chat
	mux.HandleFunc("/api/v1/chats/mensaje", protect(chats.Mensaje, hub)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/videollamada", protect(chats.Videollamada, hub)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/{user}", protect(chats.List, hub)).Methods("GET")

	// test
	mux.HandleFunc("/test", test.Test).Methods("GET")

	// websocket
	mux.HandleFunc("/ws", protect(func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		ws.ServeWs(hub, w, r, session)
	}, hub))

	mux.PathPrefix("/api/v1").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found."))
	})
	mux.PathPrefix("/").HandlerFunc(publicHandler).Methods("GET")
	return mux
}
