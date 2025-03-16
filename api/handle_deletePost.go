package api

import (
	"context"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/auth"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter) error {

	post := GetPost(ctx) // get post from context

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		return ErrorPostForbidden
	}

	// Persist changes
	err := GetPersistence(ctx).Delete(ctx, post.Id)
	if err != nil {
		log.Println("p.Delete:", err)
		return ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
