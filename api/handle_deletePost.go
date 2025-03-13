package api

import (
	"context"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/auth"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	postId := r.PathValue("postId")
	p := GetPersistence(ctx)

	post, err := p.Get(ctx, postId)
	if err != nil {
		log.Println("p.Get", err)
		return ErrorPersistenceRead
	}
	if post == nil {
		return ErrorPostNotFound
	}

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		return ErrorPostForbidden
	}

	err = p.Delete(ctx, postId)
	if err != nil {
		log.Println("p.Delete:", err)
		return ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
