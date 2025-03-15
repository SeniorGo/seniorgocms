package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SeniorGo/seniorgocms/api"
	"github.com/SeniorGo/seniorgocms/discord"
	"github.com/SeniorGo/seniorgocms/persistence"
)

var VERSION = "dev"
var DESCRIPTION = "new feature"

type Config struct {
	Addr        string                `json:"addr"`
	ServiceName string                `json:"service_name"`
	StaticsDir  string                `json:"statics_dir"`
	DataDir     string                `json:"data_dir"`
	Discord     discord.DiscordConfig `json:"discord"`
}

func main() {

	// Default config
	c := &Config{
		Addr:        ":8080",
		ServiceName: "SeniorGo - Latam",
		DataDir:     "./data",
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

	p, err := persistence.NewInDisk[api.Post](c.DataDir)
	if err != nil {
		log.Println("Error creating persistence file:", err.Error())
		return
	}

	// Instanciamos API y server
	m := api.NewApi(VERSION, c.StaticsDir, p)
	s := http.Server{
		Addr:    c.Addr,
		Handler: m,
	}

	// Start server
	log.Println("Listening on", s.Addr)
	err = s.ListenAndServe() // this call is blocking
	if err != nil {
		log.Fatal(err)
	}
}
