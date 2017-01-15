package helper

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func PuedoVer(relacion int8, permiso string) bool {
	if permiso == "private" {
		return false
	}
	if permiso == "public" {
		return true
	}
	if relacion == 1 {
		return true
	}
	return false
}

// getLetterRandom genera una letra aleatoria
func getLetterRandom() string {
	var r = rand.Intn(60)
	var b = make([]byte, 1)
	if r < 10 {
		b[0] = byte(r + 48)
	} else if r < 35 {
		b[0] = byte(r + 55)

	} else {
		b[0] = byte(r + 62)
	}
	return string(b)
}

// GenerateRandomString genera un conjunto de caracteres aleatorios
func GenerateRandomString(s int) string {
	var r = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(time.Now().Unix()))))
	r = strings.Replace(r, "==", "", 1)
	l := len(r)
	for i := 0; i < s-l; i++ {
		r += getLetterRandom()
	}
	return r
}

// Encript genera un hash aleatorio de la contraseña
func Encript(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// IsValid Verifica si la contraseña y el hash corresponden, es decir, si esa es la contraseña en bd
func IsValid(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// IsValidPermision valida si el permiso mencionado corresponde a uno de los permitidos
func IsValidPermision(permiso string) bool {
	return permiso == "private" || permiso == "friends" || permiso == "public"
}
