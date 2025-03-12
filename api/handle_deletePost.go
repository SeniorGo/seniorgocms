package api

import (
	"context"
	"log"
	"net/http"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	postId := r.PathValue("postId")
	p := GetPersistence(ctx)

	err := p.Delete(ctx, postId)
	if err != nil {
		log.Println("p.Delete:", err)
		return ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
