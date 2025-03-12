package main

import (
	"log"
	"net/http"
)

func main() {

	s := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL)

			if r.URL.Path == "/" {
				w.Write([]byte("<h1>Inicio</h1>"))
				return
			}

			if r.URL.Path == "/login" {
				w.Write([]byte("<h1>Login</h1>"))
				return
			}

			// ...
		}),
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
