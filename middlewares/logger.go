package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/juliotorresmoreno/unravel-server/helper"
)

func Cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			helper.Cors(w, r)
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

type logger struct {
	*data
	http.ResponseWriter
}

type data struct {
	statusCode int
}

func (el logger) Write(p []byte) (int, error) {
	return el.ResponseWriter.Write(p)
}
func (el logger) Header() http.Header {
	return el.ResponseWriter.Header()
}
func (el logger) WriteHeader(statusCode int) {
	el.data.statusCode = statusCode
	el.ResponseWriter.WriteHeader(statusCode)
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		log := logger{
			data:           &data{},
			ResponseWriter: w,
		}
		handler.ServeHTTP(log, r)
		t2 := time.Now()
		t3 := int(t2.UnixNano()-t1.UnixNano()) / (1000 * 1000)

		var d *color.Color
		if log.data.statusCode >= 400 {
			d = color.New(color.FgRed, color.Bold)
		} else if log.data.statusCode >= 300 {
			d = color.New(color.FgYellow, color.Bold)
		} else if log.data.statusCode >= 200 {
			d = color.New(color.FgGreen, color.Bold)
		} else if log.data.statusCode >= 100 {
			d = color.New(color.FgBlue, color.Bold)
		}
		d.Printf("%v ", log.data.statusCode)
		fmt.Println(r.Method, r.URL.Path, t3, "ms")
	})
}
