package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SeniorGo/seniorgocms/api"
)

var VERSION = "dev"

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

	fmt.Println(c.ServiceName, VERSION)

	fmt.Println("config:", c)

	m := api.NewApi(VERSION)

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
