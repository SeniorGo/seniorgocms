package api

import (
	"net/http"
	"time"
)

type Post struct {
	Id string `json:"id"`

	Title string `json:"title"`
	Body  string `json:"body"`

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

	return nil
}
