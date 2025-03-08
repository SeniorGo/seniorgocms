package api

import (
	"net/http"

	"github.com/fulldump/box"
)

func NewApi(version string) http.Handler {

	b := box.NewBox()

	b.HandleFunc("GET", "/", HandleHome)
	b.HandleFunc("GET", "/login", HandleLogin)
	b.HandleFunc("POST", "/hello", HandleHello)
	b.HandleFunc("GET", "/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	return b
}
