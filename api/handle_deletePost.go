package api

import (
	"context"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
)

type deletePost struct {
	postRepo persistence.Persistencer[Post]
}

func newDeletePost(postRepo persistence.Persistencer[Post]) *deletePost {
	return &deletePost{postRepo: postRepo}
}

func (h *deletePost) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	postId := r.PathValue("postId")

	post, err := h.postRepo.Get(ctx, postId)
	if err != nil {
		log.Println("h.postRepo.Get", err)
		return ErrorPersistenceRead
	}
	if post == nil {
		return ErrorPostNotFound
	}

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		return ErrorPostForbidden
	}

	err = h.postRepo.Delete(ctx, postId)
	if err != nil {
		log.Println("h.postRepo.Delete:", err)
		return ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
