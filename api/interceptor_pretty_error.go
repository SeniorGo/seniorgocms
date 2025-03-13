package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fulldump/box"
)

func PrettyError(next box.H) box.H {
	return func(ctx context.Context) {
		next(ctx)

		err := box.GetError(ctx)
		if err != nil {

			httpErr, ok := err.(HttpError)
			if !ok {
				httpErr = HttpError{
					Status:      http.StatusInternalServerError,
					Title:       "Unexpected error",
					Description: err.Error(),
				}
			}
			if httpErr.Title == "" {
				httpErr.Title = http.StatusText(httpErr.Status)
			}
			if httpErr.Title == "" {
				httpErr.Title = "ERROR"
			}

			w := box.GetResponse(ctx)
			w.Header().Set("X-Robots-Tag", "noindex,nofollow")
			w.WriteHeader(httpErr.Status)

			r := box.GetRequest(ctx)

			if strings.Contains(r.Header.Get("Accept"), "text/html") {
				w.Write([]byte("<h1>ERROR: " + httpErr.Title + "</h1>"))
				w.Write([]byte("<p>ERROR: " + httpErr.Description + "</p>"))
			} else {
				json.NewEncoder(w).Encode(map[string]any{
					"error": httpErr,
				})
			}
		}
	}
}
