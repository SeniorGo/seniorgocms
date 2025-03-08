package api

import (
	"net/http"
)

func NewApi(version string) http.Handler {

	m := http.NewServeMux()

	m.HandleFunc("GET /", HandleHome)
	m.HandleFunc("GET /login", HandleLogin)
	m.HandleFunc("POST /hello", HandleHello)
	m.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	return m
}
