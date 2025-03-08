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

func HandleHello(payload *HelloRequest) any {

	// Sanitize
	payload.Name = strings.TrimSpace(payload.Name)

	// Validate
	if payload.Name == "" {
		return HttpError{
			Status:      http.StatusBadRequest,
			Description: "El campo 'name' es obligatorio y no puede estar vacío",
		}
	}

	// Send response
	return HelloResponse{
		Message: "Hello " + payload.Name + "!",
	}
}
