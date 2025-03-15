package api

import (
	"context"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/persistence"
)

type getPost struct {
	postRepo persistence.Persistencer[Post]
}

func newGetPost(postRepo persistence.Persistencer[Post]) *getPost {
	return &getPost{postRepo: postRepo}
}

func (h *getPost) Handle(ctx context.Context, r *http.Request) (*Post, error) {
	postId := r.PathValue("postId")

	post, err := h.postRepo.Get(ctx, postId)
	if err != nil {
		log.Println("h.postRepo.Get:", err)
		return nil, ErrorPersistenceRead
	}
	if post == nil {
		return nil, ErrorPostNotFound
	}

	return &post.Item, nil
}
