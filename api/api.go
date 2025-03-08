package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func NewApi(version string) http.Handler {

	m := http.NewServeMux()

	m.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<h1>Hello World!!</h1>
			<p>Página de inicio</p>
			<a href="/login">Login</a>


			<div id="version" style="position: absolute; left: 0; bottom: 0;"></div>
			<script>
				fetch('/version')
					.then(req => req.text())
					.then(version => {
						document.getElementById("version").innerText = version;
					});
			</script>
		`))
	})

	m.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<h1>Entra al CMS!!</h1>
			<p>Página de login</p>
			Usuario: <input type="text" name="username" /><br>
			Contraseña: <input type="text" name="password" /><br>
			<button>Entrar</button>
		`))
	})

	m.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {

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
	})

	m.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	return m
}
