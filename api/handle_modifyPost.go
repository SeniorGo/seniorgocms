package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ModifyPostRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func HandleModifyPost(ctx context.Context, r *http.Request, input *ModifyPostRequest) (*Post, error) {

	postId := r.PathValue("postId")
	p := GetPersistence(ctx)

	post, err := p.Get(ctx, postId)
	if err != nil {
		log.Println("p.Get:", err)
		return nil, ErrorPersistenceRead
	}

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

	err = p.Put(ctx, post)
	if err != nil {
		log.Println("p.Put:", err)
		return nil, ErrorPersistenceWrite
	}

	return &post.Item, nil
}
