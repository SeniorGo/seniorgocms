package api

import (
	"context"
	"github.com/SeniorGo/seniorgocms/logger"
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

	l := logger.GetLog(ctx)

	post := Post{
		Id:           uuid.NewString(),
		Author:       auth.GetAuth(ctx).User,
		Title:        input.Title,
		Body:         input.Body,
		CreationTime: time.Now(),
	}
	post.ModificationTime = post.CreationTime

	l.Info("Creando un post nuevo", "title", input.Title, "id", post.Id)

	err := post.Validate()
	if err != nil {
		l.Warn("El post no es válido: " + err.Error())
		return nil, err
	}

	p := GetPostPersistence(ctx)
	err = p.Put(context.Background(), &persistence.ItemWithId[Post]{
		Id:   post.Id,
		Item: post,
	})
	if err != nil {
		l.Error("p.Put: " + err.Error())
		return nil, ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusCreated)

	l.Info("Post creado correctamente")

	return &post, nil
}
