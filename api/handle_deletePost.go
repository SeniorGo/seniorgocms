package api

import (
	"context"
	"net/http"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")
	p := GetPersistence(ctx)

	err := p.Delete(ctx, postId)
	if err != nil {
		// TODO: handle error
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
