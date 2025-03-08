package api

import (
	"net/http"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/biff"
)

func TestNewApi(t *testing.T) {

	h := NewApi("testversion", "")
	a := apitest.NewWithHandler(h)

	t.Run("Request /version", func(t *testing.T) {
		r := a.Request("GET", "/version").Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(r.BodyString(), "testversion")
	})

	t.Run("Request /hello", func(t *testing.T) {
		r := a.Request("POST", "/hello").
			WithBodyJson(map[string]string{"name": "Manu"}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(r.BodyJsonMap()["message"], "Hello Manu!")
	})

}
