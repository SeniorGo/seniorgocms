package api

import (
	"net/http"

	"github.com/fulldump/box"

	"github.com/SeniorGo/seniorgocms/statics"
)

func NewApi(version, staticsDir string) http.Handler {

	b := box.NewBox()

	b.WithInterceptors(PrettyError)

	b.HandleResourceNotFound = HandleNotFound
	b.HandleMethodNotAllowed = HandleMethodNotAllowed

	b.Handle("GET", "/bad", HandleBad)
	b.Handle("POST", "/hello", HandleHello)
	b.Handle("GET", "/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	// Mount statics
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).WithName("serveStatics")

	return b
}
