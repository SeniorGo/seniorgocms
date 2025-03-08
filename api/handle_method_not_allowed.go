package api

import (
	"context"
	"net/http"

	"github.com/fulldump/box"
)

func HandleMethodNotAllowed(ctx context.Context, r *http.Request) {

	box.SetError(ctx, HttpError{
		Status:      http.StatusMethodNotAllowed,
		Description: "This method '" + r.Method + "' is not allowed",
	})
}
