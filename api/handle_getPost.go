package api

import (
	"context"
	"log"
	"net/http"
)

func HandleGetPost(ctx context.Context, r *http.Request) (*Post, error) {

	postId := r.PathValue("postId")
	p := GetPersistence(ctx)

	post, err := p.Get(ctx, postId)
	if err != nil {
		log.Println("p.Get:", err)
		return nil, HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Problem reading from persistence layer",
		}
	}

	return &post.Item, nil
}
