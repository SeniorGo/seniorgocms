package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
)

type modifyPost struct {
	postRepo persistence.Persistencer[Post]
}

type ModifyPostRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func newModifyPost(postRepo persistence.Persistencer[Post]) *modifyPost {
	return &modifyPost{postRepo: postRepo}
}

func (h *modifyPost) Handle(ctx context.Context, r *http.Request, input *ModifyPostRequest) (*Post, error) {
	postId := r.PathValue("postId")

	post, err := h.postRepo.Get(ctx, postId)
	if err != nil {
		log.Println("h.postRepo.Get:", err)
		return nil, ErrorPersistenceRead
	}
	if post == nil {
		return nil, ErrorPostNotFound
	}

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		return nil, ErrorPostForbidden
	}

	post.Item.Author = auth.GetAuth(ctx).User // Update user data
	post.Item.ModificationTime = time.Now()

	if input.Title != nil {
		post.Item.Title = *input.Title
	}

	if input.Body != nil {
		post.Item.Body = *input.Body
	}

	err = post.Item.Validate()
	if err != nil {
		return nil, err
	}

	err = h.postRepo.Put(ctx, post)
	if err != nil {
		log.Println("h.postRepo.Put:", err)
		return nil, ErrorPersistenceWrite
	}

	return &post.Item, nil
}
