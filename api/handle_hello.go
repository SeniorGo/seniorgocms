package api

import (
	"net/http"
	"strings"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func HandleHello(payload *HelloRequest) (*HelloResponse, error) {

	// Sanitize
	payload.Name = strings.TrimSpace(payload.Name)

	// Validate
	if payload.Name == "" {
		return nil, HttpError{
			Status:      http.StatusBadRequest,
			Description: "El campo 'name' es obligatorio y no puede estar vacío",
		}
	}

	// Send response
	return &HelloResponse{
		Message: "Hello " + payload.Name + "!",
	}, nil
}
