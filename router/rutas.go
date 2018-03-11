package router

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/controllers/experience"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/controllers/auth"
	"github.com/juliotorresmoreno/unravel-server/controllers/auth/oauth"
	"github.com/juliotorresmoreno/unravel-server/controllers/category"
	"github.com/juliotorresmoreno/unravel-server/controllers/chats"
	"github.com/juliotorresmoreno/unravel-server/controllers/friends"
	"github.com/juliotorresmoreno/unravel-server/controllers/galery"
	"github.com/juliotorresmoreno/unravel-server/controllers/groups"
	"github.com/juliotorresmoreno/unravel-server/controllers/news"
	"github.com/juliotorresmoreno/unravel-server/controllers/profile"
	"github.com/juliotorresmoreno/unravel-server/controllers/users"
	api "github.com/juliotorresmoreno/unravel-server/graphql"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
	"github.com/juliotorresmoreno/unravel-server/test"
	"github.com/juliotorresmoreno/unravel-server/ws"
	"github.com/nytimes/gziphandler"
)

// GetHandler aca se establecen las rutas del router
func GetHandler() http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)
	var hub = ws.GetHub()

	mux.PathPrefix("/").HandlerFunc(helper.HandleCors).Methods("OPTIONS")

	//graphql
	mux.HandleFunc("/api/v2/graphql", protect(func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		result := api.ExecuteQuery(r.URL.Query()["query"][0])
		json.NewEncoder(w).Encode(result)
	}, hub, true))

	// auth
	mux.HandleFunc("/api/v1/auth/registrar", auth.Registrar).Methods("POST")
	mux.HandleFunc("/api/v1/auth/login", auth.Login).Methods("POST")
	mux.HandleFunc("/api/v1/auth/session", protect(auth.Session, hub, false)).Methods("GET")
	mux.HandleFunc("/api/v1/auth/logout", auth.Logout).Methods("GET")
	mux.HandleFunc("/api/v1/auth/recovery", auth.Recovery).Methods("POST")
	mux.HandleFunc("/api/v1/auth/password", auth.Password).Methods("POST")
	mux.HandleFunc("/api/v1/auth/password_change", protect(auth.PasswordChange, hub, true)).Methods("POST")
	mux.HandleFunc("/oauth2callback", auth.Oauth2Callback).Methods("GET")
	mux.HandleFunc("/auth/facebook", oauth.HandleFacebook).Methods("GET")
	mux.HandleFunc("/auth/github", oauth.HandleGithub).Methods("GET")
	mux.HandleFunc("/auth/google", oauth.HandleGoogle).Methods("GET")

	// profile
	mux.HandleFunc("/api/v1/profile", protect(profile.Profile, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/profile/{user}", protect(profile.Profile, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/profile", protect(profile.Update, hub, true)).Methods("POST", "PUT", "OPTIONS")

	// friends
	mux.HandleFunc("/api/v1/friends", protect(friends.ListFriends, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/friends/{usuario}", protect(friends.ListFriends, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/friends/add", protect(friends.Add, hub, true)).Methods("POST", "PUT")
	mux.HandleFunc("/api/v1/friends/reject", protect(friends.RejectFriend, hub, true)).Methods("POST", "DELETE")

	// users
	mux.HandleFunc("/api/v1/users", protect(users.Find, hub, true)).Methods("GET")

	// news
	mux.HandleFunc("/api/v1/news", protect(news.GetNews, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/news/public", protect(news.Publicar, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/news/like", protect(news.Like, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/news/comentar", protect(news.Comentar, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/{usuario}/news", protect(news.GetNews, hub, true)).Methods("GET")

	// galery
	mux.HandleFunc("/api/v1/galery", protect(galery.ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery", protect(galery.Save, hub, true)).Methods("POST")

	mux.HandleFunc("/api/v1/galery/delete", protect(galery.EliminarImagen, hub, true)).Methods("POST", "DELETE")

	mux.HandleFunc("/api/v1/galery/fotoPerfil", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/upload", protect(galery.Upload, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/galery/fotoPerfil", protect(galery.SetFotoPerfil, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/galery/fotoPerfil/{usuario}", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}/describe", protect(galery.DescribeGaleria, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}/preview", protect(galery.ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}/{imagen}", protect(galery.ViewImagen, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/galery/{galery}", protect(galery.ListarImagenes, hub, true)).Methods("GET")

	mux.HandleFunc("/api/v1/{usuario}/galery", protect(galery.ListarGalerias, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/fotoPerfil", protect(galery.GetFotoPerfil, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}", protect(galery.ListarImagenes, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}/preview", protect(galery.ViewPreview, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/{usuario}/galery/{galery}/{imagen}", protect(galery.ViewImagen, hub, true)).Methods("GET")

	// groups
	mux.HandleFunc("/api/v1/groups", protect(groups.ObtenerGrupos, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/groups/all", protect(groups.ObtenerTodosGrupos, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/groups", protect(groups.Save, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/groups/changePreview", protect(groups.ChangePreview, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/groups/{group}", protect(groups.Describe, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/groups/{group}/preview", protect(groups.Preview, hub, true)).Methods("GET")

	// category
	mux.HandleFunc("/api/v1/category", protect(category.GetCategorys, hub, true)).Methods("GET")

	// chat
	mux.HandleFunc("/api/v1/chats/mensaje", protect(chats.Mensaje, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/videollamada", protect(chats.VideoLlamada, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/rechazarvideollamada", protect(chats.RechazarVideoLlamada, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/chats/{user}", protect(chats.GetConversacion, hub, true)).Methods("GET")

	mux.PathPrefix("/api/v1/experience").
		Handler(helper.StripPrefix(
			"/api/v1/experience",
			experience.NewRouter(hub),
		))

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

	// archivos estaticos
	withoutGz := http.HandlerFunc(publicHandler)
	withGz := gziphandler.GzipHandler(withoutGz)
	mux.PathPrefix("/").Handler(withGz).Methods("GET")

	return mux
}
