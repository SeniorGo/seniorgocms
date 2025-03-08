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

	s := http.Server{
		Addr: c.Addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.String())

			if r.URL.Path == "/" {
				w.Write([]byte(`
				<h1>Hello World!!</h1>
				<p>Página de inicio</p>
				<a href="/login">Login</a>
				`))
				return
			}

			if r.URL.Path == "/login" {
				w.Write([]byte(`
				<h1>Entra al CMS!!</h1>
				<p>Página de login</p>
				Usuario: <input type="text" name="username" /><br>
				Contraseña: <input type="text" name="password" /><br>		
				<button>Entrar</button>
				`))
				return
			}

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`
				<h1>NOT FOUND!!</h1>
				`))
			return

		}),
	}

	log.Println("Listening on", s.Addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
