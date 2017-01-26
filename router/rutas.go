package router

import (
	"net/http"

	"../controllers/auth"
	"../controllers/chats"
	"../controllers/friends"
	"../controllers/galery"
	"../controllers/profile"
	"../controllers/users"
	"../models"
	"../test"
	"../ws"
	"github.com/gorilla/mux"
)

// GetHandler aca se establecen las rutas del router
func GetHandler() http.Handler {
	var mux = mux.NewRouter().StrictSlash(false)
	var hub = ws.GetHub()

	// auth
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")
	mux.HandleFunc("/api/v1/auth/login", auth.Login).Methods("POST")
	mux.HandleFunc("/api/v1/auth/session", protect(auth.Session, hub, false)).Methods("GET")
	mux.HandleFunc("/api/v1/auth/logout", auth.Logout).Methods("GET")

	// profile
	mux.HandleFunc("/api/v1/profile", protect(profile.Profile, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/profile/{user}", protect(profile.Profile, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/profile", protect(profile.Update, hub, true)).Methods("POST", "PUT")

	// friends
	mux.HandleFunc("/api/v1/friends", protect(friends.ListFriends, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/friends/add", protect(friends.Add, hub, true)).Methods("POST", "PUT")
	mux.HandleFunc("/api/v1/friends/reject", protect(friends.RejectFriend, hub, true)).Methods("POST", "DELETE")

	// users
	mux.HandleFunc("/api/v1/users", protect(users.Find, hub, true)).Methods("GET")

	// galery
	mux.HandleFunc("/api/v1/galery/fotoPerfil", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery", protect(galery.ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/create", protect(galery.Create, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/galery/upload", protect(galery.Upload, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/galery/fotoPerfil", protect(galery.SetFotoPerfil, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/galery/fotoPerfil/{usuario}", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}/preview", protect(galery.ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}/{imagen}", protect(galery.ViewImagen, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}", protect(galery.ListarImagenes, hub, true)).Methods("GET")

	mux.HandleFunc("/api/v1/{usuario}/galery", protect(galery.ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/fotoPerfil", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}", protect(galery.ListarImagenes, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}/preview", protect(galery.ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}/{imagen}", protect(galery.ViewImagen, hub, true)).Methods("GET")

	// chat
	mux.HandleFunc("/api/v1/chats/mensaje", protect(chats.Mensaje, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/videollamada", protect(chats.Videollamada, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/{user}", protect(chats.List, hub, true)).Methods("GET")

	// test
	mux.HandleFunc("/test", test.Test).Methods("GET")

	// websocket
	mux.HandleFunc("/ws", protect(func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		ws.ServeWs(hub, w, r, session)
	}, hub, true))

	mux.PathPrefix("/api/v1").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found."))
	})

	//http://localhost:8080/api/v1/jtorres990//galery/Nombre?token=MTQ4NTM5NTAxMwgRmy1IPKG0EVhTSEq5HQa68xR77SUF1nmB9vwVk6DAY3DRx4JXw19efP8IiaBUkbl8g2NQib2dRtciGK3XcDgJ
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}", protect(galery.ListarImagenes, hub, true)).Methods("GET")
	mux.PathPrefix("/").HandlerFunc(publicHandler).Methods("GET")
	return mux

}
