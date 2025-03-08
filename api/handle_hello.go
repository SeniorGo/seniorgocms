package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		Name string `json:"name"`
	}{}

	// Read payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El JSON que me has enviado no es válido"))
		return
	}

	// Sanitize
	payload.Name = strings.TrimSpace(payload.Name)

	// Validate
	if payload.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El campo 'name' es obligatorio y no puede estar vacío"))
		return
	}

	// Send response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello " + payload.Name + "!",
	})
}
