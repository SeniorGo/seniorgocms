package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
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
	Log         ConfigLog             `json:"log"`
}

type ConfigLog struct {
	Type  string `json:"type"`
	Level string `json:"level"` // TODO: hacer luego
}

func main() {

	// Default config
	c := &Config{
		Addr:        ":8080",
		ServiceName: "SeniorGo - Latam",
		DataDir:     "./data",
		Log: ConfigLog{
			Type: "json",
		},
	}

	// Read config
	f, err := os.Open("./config.json")
	if err == nil {
		json.NewDecoder(f).Decode(&c)
	}
	fmt.Println(c.ServiceName, VERSION)

	var logHandler slog.Handler
	if c.Log.Type == "text" {
		logHandler = slog.NewTextHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	}
	l := slog.New(logHandler).With("version", VERSION)

	// Notify to discord
	msg := c.ServiceName + ": Nueva version " + VERSION + "\n" + DESCRIPTION
	l.Info(msg)
	err = discord.Notify(c.Discord, msg)
	if err != nil {
		log.Println("Error sending notification:", err.Error())
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
	m := api.NewApi(VERSION, c.StaticsDir, postPersistencer, categoryPersistencer, l)
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
