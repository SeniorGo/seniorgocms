package api

import (
	"net/http"
	"time"
)

type Category struct {
	Id string `json:"id"`

	Name string `json:"name"`

	CreationTime     time.Time `json:"creation_time"`
	ModificationTime time.Time `json:"modification_time"`
}

func (c *Category) Validate() error {
	if len(c.Name) > 1024 {
		return HttpError{
			Status:      http.StatusBadRequest,
			Description: "name is too long",
		}
	}

	return nil
}
