package api

import (
	"net/http"

	"github.com/fulldump/box"

	"github.com/SeniorGo/seniorgocms/statics"
)

func NewApi(version, staticsDir string) http.Handler {

	b := box.NewBox()

	b.HandleResourceNotFound = HandleNotFound

	b.Handle("POST", "/hello", HandleHello)
	b.HandleFunc("GET", "/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	// Mount statics
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).WithName("serveStatics")

	return b
}
