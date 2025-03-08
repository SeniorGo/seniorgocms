package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type DiscordConfig struct {
	Authorization   string `json:"authorization"`
	SuperProperties string `json:"super_properties"`
}

func Notify(config DiscordConfig, content string) error {

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
