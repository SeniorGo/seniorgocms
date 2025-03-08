package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SeniorGo/seniorgocms/api"
	"github.com/SeniorGo/seniorgocms/discord"
)

var VERSION = "dev"
var DESCRIPTION = "new feature"

type Config struct {
	Addr        string                `json:"addr"`
	ServiceName string                `json:"service_name"`
	Discord     discord.DiscordConfig `json:"discord"`
}

func main() {

	// Default config
	c := &Config{
		Addr:        ":8080",
		ServiceName: "SeniorGo - Latam",
	}

	// Read config
	f, err := os.Open("./config.json")
	if err == nil {
		json.NewDecoder(f).Decode(&c)
	}
	fmt.Println(c.ServiceName, VERSION)

	// Notify to discord
	msg := c.ServiceName + ": Nueva version " + VERSION + "\n" + DESCRIPTION
	log.Println(msg)
	err = discord.Notify(c.Discord, msg)
	if err != nil {
		log.Println("Error sending notification:", err.Error())
	}

	// Instanciamos API y server
	m := api.NewApi(VERSION)
	s := http.Server{
		Addr:    c.Addr,
		Handler: api.MiddlewareAccessLog(m.ServeHTTP),
	}

	// Start server
	log.Println("Listening on", s.Addr)
	err = s.ListenAndServe() // this call is blocking
	if err != nil {
		log.Fatal(err)
	}
}
