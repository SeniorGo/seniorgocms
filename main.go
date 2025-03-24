package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/SeniorGo/seniorgocms/api"
	"github.com/SeniorGo/seniorgocms/discord"
	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/utils"
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
	// Initialize global logger
	utils.InitLogger()
	logger := utils.GlobalLogger

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
	logger.Info("Iniciando %s versión %s", c.ServiceName, VERSION)

	// Notify to discord
	msg := c.ServiceName + ": Nueva version " + VERSION + "\n" + DESCRIPTION
	logger.Info(msg)
	err = discord.Notify(c.Discord, msg)
	if err != nil {
		logger.Error("Error sending notification: %v", err)
	}

	postPersistencer, err := persistence.NewInDisk[api.Post](c.DataDir + "/posts")
	if err != nil {
		log.Println("Error creating persistence file:", err.Error())
		return
	}

	categoryPersistencer, err := persistence.NewInDisk[api.Category](c.DataDir + "/categories")
	if err != nil {
		log.Println("Error creating persistence file:", err.Error())
		return
	}

	// Instanciamos API y server
	m := api.NewApi(VERSION, c.StaticsDir, postPersistencer, categoryPersistencer)
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
