package api

import (
	"context"
	"github.com/SeniorGo/seniorgocms/logger"
	"net/http"
	"time"

	"github.com/SeniorGo/seniorgocms/auth"
)

type ModifyPostRequest struct {
	Title *string   `json:"title"`
	Body  *string   `json:"body"`
	Tags  *[]string `json:"tags"`
}

func HandleModifyPost(ctx context.Context, r *http.Request, input *ModifyPostRequest) (*Post, error) {

	l := logger.GetLog(ctx)

	postId := r.PathValue("postId")
	p := GetPostPersistence(ctx)

	l.Info("Modificando el post", "title", input.Title, "postId", postId)

	post, err := p.Get(ctx, postId)
	if err != nil {
		l.Error("Error getting post", "error", err.Error())
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

	if input.Tags != nil {
		post.Item.Tags = *input.Tags
	}

	err = post.Item.Validate()
	if err != nil {
		return nil, err
	}

	err = p.Put(ctx, post)
	if err != nil {
		l.Error("Error updating post", "error", err.Error())
		return nil, ErrorPersistenceWrite
	}

	l.Info("Post modificado correctamente")

	return &post.Item, nil
}
