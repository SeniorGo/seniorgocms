package api

import (
	"context"
	"github.com/SeniorGo/seniorgocms/logger"
	"net/http"

	"github.com/SeniorGo/seniorgocms/auth"
)

func HandleDeletePost(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	l := logger.GetLog(ctx)

	postId := r.PathValue("postId")
	p := GetPostPersistence(ctx)

	l.Info("Eliminando el post", "postId", postId)

	post, err := p.Get(ctx, postId)
	if err != nil {
		l.Error("p.Get", err)
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
		l.Error("p.Delete:", err)
		return ErrorPersistenceWrite
	}

	l.Info("Post eliminado correctamente")

	w.WriteHeader(http.StatusNoContent)
	return nil
}
