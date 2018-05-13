package router

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/unravel-server/controllers/educacion"
	"github.com/juliotorresmoreno/unravel-server/controllers/experience"
	"github.com/juliotorresmoreno/unravel-server/controllers/skill"
	"github.com/juliotorresmoreno/unravel-server/middlewares"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/controllers/auth"
	"github.com/juliotorresmoreno/unravel-server/controllers/category"
	"github.com/juliotorresmoreno/unravel-server/controllers/chats"
	"github.com/juliotorresmoreno/unravel-server/controllers/friends"
	"github.com/juliotorresmoreno/unravel-server/controllers/galery"
	"github.com/juliotorresmoreno/unravel-server/controllers/geo"
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

// NewRouter aca se establecen las rutas del router
func NewRouter() http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)
	var hub = ws.GetHub()

	mux.Use(middlewares.Logger)
	mux.Use(middlewares.Cors)
	mux.Use(gziphandler.GzipHandler)

	//graphql
	mux.HandleFunc("/api/v2/graphql", protect(func(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
		result := api.ExecuteQuery(r.URL.Query()["query"][0])
		json.NewEncoder(w).Encode(result)
	}, hub, true))

	// auth
	mux.PathPrefix("/api/v1/auth").
		Handler(helper.StripPrefix(
			"/api/v1/auth",
			auth.NewRouter(hub),
		))

	// profile
	mux.PathPrefix("/api/v1/profile").
		Handler(helper.StripPrefix(
			"/api/v1/profile",
			profile.NewRouter(hub),
		))

	// friends
	mux.PathPrefix("/api/v1/friends").
		Handler(helper.StripPrefix(
			"/api/v1/friends",
			friends.NewRouter(hub),
		))

	// users
	mux.HandleFunc("/api/v1/users", protect(users.Find, hub, true)).Methods("GET")

	// news
	mux.HandleFunc("/api/v1/news", protect(news.GetNews, hub, true)).Methods("GET")
	mux.HandleFunc("/api/v1/news/public", protect(news.Publicar, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/news/like", protect(news.Like, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/news/comentar", protect(news.Comentar, hub, true)).Methods("POST")
	mux.HandleFunc("/api/v1/{usuario}/news", protect(news.GetNews, hub, true)).Methods("GET")

	// galery
	mux.PathPrefix("/api/v1/galery").
		Handler(helper.StripPrefix(
			"/api/v1/galery",
			galery.NewRouter(hub),
		))
	mux.PathPrefix("/api/v1/{usuario}/galery").
		Handler(helper.StripPrefix(
			"/api/v1",
			galery.NewUserRouter(hub),
		))

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
	mux.PathPrefix("/api/v1/chats").
		Handler(helper.StripPrefix(
			"/api/v1/chats",
			chats.NewRouter(hub),
		))

	mux.PathPrefix("/api/v1/experience").
		Handler(helper.StripPrefix(
			"/api/v1/experience",
			experience.NewRouter(hub),
		))

	mux.PathPrefix("/api/v1/educacion").
		Handler(helper.StripPrefix(
			"/api/v1/educacion",
			educacion.NewRouter(hub),
		))

	mux.PathPrefix("/api/v1/skills").
		Handler(helper.StripPrefix(
			"/api/v1/skills",
			skill.NewRouter(hub),
		))

	mux.PathPrefix("/api/v1/geo").
		Handler(helper.StripPrefix(
			"/api/v1/geo",
			geo.NewRouter(hub),
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
	mux.PathPrefix("/").HandlerFunc(publicHandler).Methods("GET")

	return mux
}
