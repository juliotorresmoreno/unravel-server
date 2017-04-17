package auth

import (
	"net/http"
	"time"

	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
)

// Logout cerrar session
func Logout(w http.ResponseWriter, r *http.Request) {
	var cache = models.GetCache()
	var token string = helper.GetToken(r)
	if token != "" {
		cache.Del(token)
		http.SetCookie(w, &http.Cookie{
			MaxAge:   0,
			Secure:   false,
			HttpOnly: true,
			Name:     "token",
			Value:    "",
			Path:     "/",
			Expires:  time.Now(),
		})
	}
	helper.Cors(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":true}"))
}
