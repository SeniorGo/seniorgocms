package api

import (
	"net/http"
	"time"

	"github.com/SeniorGo/seniorgocms/auth"
)

type Post struct {
	Id string `json:"id"`

	Author auth.User `json:"author"`

	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`

	CreationTime     time.Time `json:"creation_time"`
	ModificationTime time.Time `json:"modification_time"`
}

func (e *Post) Validate() error {

	if len(e.Title) > 1024 {
		return HttpError{
			Status:      http.StatusBadRequest,
			Description: "title is too long",
		}
	}

	if len(e.Body) > 100*1024 {
		return HttpError{
			Status:      http.StatusBadRequest,
			Description: "body too long",
		}
	}

	for _, tag := range e.Tags {
		if len(tag) > 128 {
			return HttpError{
				Status:      http.StatusBadRequest,
				Description: "tag is too long (max 128 characters)",
			}
		}
	}

	return nil
}
