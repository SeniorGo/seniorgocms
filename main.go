package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/SeniorGo/seniorgocms/api"
)

var VERSION = "dev"
var DESCRIPTION = "new feature"

type Config struct {
	Addr        string        `json:"addr"`
	ServiceName string        `json:"service_name"`
	Discord     DiscordConfig `json:"discord"`
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

	msg := c.ServiceName + ": Nueva version " + VERSION + "\n" + DESCRIPTION
	log.Println(msg)
	err = NotifyDiscord(c.Discord, msg)
	if err != nil {
		log.Println("Error sending notification:", err.Error())
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

type DiscordConfig struct {
	Authorization   string `json:"authorization"`
	SuperProperties string `json:"super_properties"`
}

func NotifyDiscord(config DiscordConfig, content string) error {

	payload, _ := json.Marshal(map[string]any{
		"mobile_network_type": "unknown",
		"content":             content,
		"nonce":               strconv.FormatInt(time.Now().UnixNano(), 10),
		"tts":                 false,
		"flags":               0,
	})

	u := "https://discord.com/api/v9/channels/1242312465052602438/messages"
	body := bytes.NewReader(payload)
	r, err := http.NewRequest("POST", u, body)
	if err != nil {
		return err
	}

	r.Header.Set("content-type", "application/json")
	r.Header.Set("authorization", config.Authorization)
	r.Header.Set("x-super-properties", config.SuperProperties)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Unexpected status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
