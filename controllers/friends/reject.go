package friends

import (
	"net/http"

	"github.com/unravel-server/models"
	"github.com/unravel-server/ws"
)

// RejectFriend Rechazar amistad
func RejectFriend(w http.ResponseWriter, r *http.Request, session *models.User, hub *ws.Hub) {
	models.RejectFriends(session.Usuario, r.PostFormValue("user"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": true}"))
}
