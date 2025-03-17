package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
)

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func HandleCreatePost(input *CreatePostRequest, w http.ResponseWriter, ctx context.Context) (*Post, error) {

	post := Post{
		Id:           uuid.NewString(),
		Author:       auth.GetAuth(ctx).User,
		Title:        input.Title,
		Body:         input.Body,
		CreationTime: time.Now(),
	}
	post.ModificationTime = post.CreationTime

	err := post.Validate()
	if err != nil {
		return nil, err
	}

	p := GetPostPersistence(ctx)
	err = p.Put(context.Background(), &persistence.ItemWithId[Post]{
		Id:   post.Id,
		Item: post,
	})
	if err != nil {
		log.Println("p.Put:", err)
		return nil, ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusCreated)

	return &post, nil
}
