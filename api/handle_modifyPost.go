package api

import (
	"context"
	"log"
	"time"

	"github.com/SeniorGo/seniorgocms/auth"
)

type ModifyPostRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func HandleModifyPost(ctx context.Context, input *ModifyPostRequest) (*Post, error) {

	post := GetPost(ctx) // get post from context

	// Access control
	if post.Item.Author.ID != auth.GetAuth(ctx).User.ID {
		return nil, ErrorPostForbidden
	}

	// Update fields
	post.Item.Author = auth.GetAuth(ctx).User // Update user data
	post.Item.ModificationTime = time.Now()

	if input.Title != nil {
		post.Item.Title = *input.Title
	}

	if input.Body != nil {
		post.Item.Body = *input.Body
	}

	// Validate post
	err := post.Item.Validate()
	if err != nil {
		return nil, err
	}

	// Save changes
	err = GetPersistence(ctx).Put(ctx, post)
	if err != nil {
		log.Println("p.Put:", err)
		return nil, ErrorPersistenceWrite
	}

	return &post.Item, nil
}
