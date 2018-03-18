package friends

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/unravel-server/middlewares"
	"github.com/juliotorresmoreno/unravel-server/ws"
)

func NewRouter(hub *ws.Hub) http.Handler {
	var mux = mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/", middlewares.Protect(ListFriends, hub, true)).Methods("GET")
	mux.HandleFunc("/{usuario}", middlewares.Protect(ListFriends, hub, true)).Methods("GET")
	mux.HandleFunc("/add", middlewares.Protect(Add, hub, true)).Methods("POST", "PUT")
	mux.HandleFunc("/reject", middlewares.Protect(RejectFriend, hub, true)).Methods("POST", "DELETE")

	return mux
}
