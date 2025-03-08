package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr        string `json:"addr"`
	ServiceName string `json:"service_name"`
}

func main() {

	c := &Config{
		Addr:        ":8080",
		ServiceName: "SeniorGo - Latam",
	}

	f, err := os.Open("./config.json")
	if err == nil {
		json.NewDecoder(f).Decode(&c)
	}

	fmt.Println("config:", c)

	m := http.NewServeMux()

	m.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<h1>Hello World!!</h1>
			<p>Página de inicio</p>
			<a href="/login">Login</a>
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

	s := http.Server{
		Addr:    c.Addr,
		Handler: MiddlewareAccessLog(m.ServeHTTP),
	}

	log.Println("Listening on", s.Addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func MiddlewareAccessLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	}
}
