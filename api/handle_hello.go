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

func HandleHello(payload *HelloRequest, w http.ResponseWriter) any {

	// Sanitize
	payload.Name = strings.TrimSpace(payload.Name)

	// Validate
	if payload.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El campo 'name' es obligatorio y no puede estar vacío"))
		return nil
	}

	// Send response
	return HelloResponse{
		Message: "Hello " + payload.Name + "!",
	}
}
