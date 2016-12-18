package services

import(
	"math/rand"
	"golang.org/x/crypto/bcrypt"
	"time"
	"encoding/base64"
	"strconv"
	"strings"
)

func getLetterRandom() string {
	var r int = rand.Intn(60)
	var b []byte = make([]byte, 1)
	if r < 10 {
		b[0] = byte(r + 48)
	} else if r < 35 {
		b[0] = byte(r + 55)

	} else {
		b[0] = byte(r + 62)
	}
	return string(b)
}
func GenerateRandomString(s int) (string) {
	var r string = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(time.Now().Unix()))))
	r = strings.Replace(r, "==", "", 1)
	l:= len(r)
	for i:= 0; i < s - l; i++ {
		r += getLetterRandom()
	}
	return r
}


func Encript(password string) (string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func IsValid(hash string, password string) (bool) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}