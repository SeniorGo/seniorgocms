package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/SeniorGo/seniorgocms/persistence"
)

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func HandleCreatePost(input *CreatePostRequest, w http.ResponseWriter, ctx context.Context) (*Post, error) {

	if len(input.Title) > 1024 {
		return nil, HttpError{
			Status:      http.StatusBadRequest,
			Description: "title is too long",
		}
	}
	if len(input.Body) > 100*1024 {
		return nil, errors.New("body too long")
	}

	post := Post{
		Id:           uuid.NewString(),
		Title:        input.Title,
		Body:         input.Body,
		CreationTime: time.Now(),
	}
	post.ModificationTime = post.CreationTime

	p := GetPersistence(ctx)
	err := p.Put(context.Background(), &persistence.ItemWithId[Post]{
		Id:   post.Id,
		Item: post,
	})
	if err != nil {
		log.Println("p.Put:", err)
		return nil, HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Problem writing to persistence layer",
		}
	}

	w.WriteHeader(http.StatusCreated)

	return &post, nil
}
