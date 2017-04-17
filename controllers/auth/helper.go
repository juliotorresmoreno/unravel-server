package auth

import (
	"time"

	"github.com/juliotorresmoreno/unravel-server/config"
	"github.com/juliotorresmoreno/unravel-server/helper"
	"github.com/juliotorresmoreno/unravel-server/models"
)

func autenticate(user *models.User) (string, map[string]interface{}) {
	_token := helper.GenerateRandomString(100)
	cache := models.GetCache()
	cache.Set(string(_token), user.Usuario, time.Duration(config.SESSION_DURATION)*time.Second)

	respuesta := map[string]interface{}{
		"success": true,
		"session": map[string]string{
			"usuario":   user.Usuario,
			"fullname":  user.FullName,
			"nombres":   user.Nombres,
			"apellidos": user.Apellidos,
			"token":     _token,
		},
	}

	return _token, respuesta
}
