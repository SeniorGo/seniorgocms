package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func HandleCreateCategory(input *CreateCategoryRequest, w http.ResponseWriter, ctx context.Context) (*Category, error) {
	category := Category{
		Id:           uuid.NewString(),
		Name:         input.Name,
		CreationTime: time.Now(),
	}
	category.ModificationTime = category.CreationTime

	err := category.Validate()
	if err != nil {
		return nil, err
	}

	p := GetCategoryPersistence(ctx)
	err = p.Put(context.Background(), &persistence.ItemWithId[Category]{
		Id:   category.Id,
		Item: category,
	})
	if err != nil {
		log.Println("p.Put:", err)
		return nil, ErrorPersistenceWrite
	}

	w.WriteHeader(http.StatusCreated)

	return &category, nil
}
