package helper

import (
	"net/http"
	"strings"
)

func GetCookie(r *http.Request, name string) string {
	var c string = r.Header.Get("Cookie")
	var s []string = strings.Split(c, ";")
	var t []string
	for i := 0; i < len(s); i++ {
		t = strings.Split(strings.Trim(s[i], " "), "=")
		if t[0] == name {
			return t[1]
		}
	}
	return ""
}