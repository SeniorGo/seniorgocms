package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewApi(t *testing.T) {

	h := NewApi("testversion")
	s := httptest.NewServer(h)

	t.Run("Request /version", func(t *testing.T) {
		req, _ := http.NewRequest("GET", s.URL+"/version", nil)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("response status is %d, want 200", res.StatusCode)
		}

		body, _ := ioutil.ReadAll(res.Body)
		if string(body) != "testversion" {
			t.Fatalf("response body is %s, want %s", string(body), "testversion")
		}

	})

	t.Run("Request /hello", func(t *testing.T) {
		payload := bytes.NewReader([]byte(`{"name": "Manu"}`))
		req, _ := http.NewRequest("POST", s.URL+"/hello", payload)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("response status is %d, want 200", res.StatusCode)
		}

		body := map[string]interface{}{}
		json.NewDecoder(res.Body).Decode(&body)

		have := body["message"]
		want := "Hello Manu!"
		if have != want {
			t.Fatalf("response body is %s, want %s", have, want)
		}
	})
}
