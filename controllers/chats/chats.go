package chats

import (
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("yea"))
}