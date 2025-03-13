package api

import (
	"net/http"
	"strings"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/biff"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
)

func TestNewApi(t *testing.T) {

	// p := persistence.NewInMemory[Post]()
	p, err := persistence.NewInDisk[Post](t.TempDir())
	biff.AssertNil(err)

	authHeader := `{"user":{"id":"user-test"}}`

	h := NewApi("testversion", "", p)
	a := apitest.NewWithHandler(h)

	t.Run("Request /version", func(t *testing.T) {
		r := a.Request("GET", "/version").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(r.BodyString(), "testversion")
	})

	t.Run("Request /hello", func(t *testing.T) {
		r := a.Request("POST", "/hello").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]string{"name": "Manu"}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(r.BodyJsonMap()["message"], "Hello Manu!")
	})

	t.Run("List posts (empty)", func(t *testing.T) {
		r := a.Request("GET", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)

		expectedBody := []map[string]interface{}{}
		biff.AssertEqualJson(r.BodyJson(), expectedBody)
	})

	t.Run("Create Post", func(t *testing.T) {
		r := a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]string{
				"title": "My Post",
				"body":  "This is my body",
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusCreated)

		body := r.BodyJsonMap()
		expectedBody := map[string]interface{}{
			"id":                body["id"],
			"title":             "My Post",
			"body":              "This is my body",
			"creation_time":     body["creation_time"],
			"modification_time": body["modification_time"],
		}
		biff.AssertEqual(r.BodyJsonMap(), expectedBody)
	})

	t.Run("List posts (1)", func(t *testing.T) {
		r := a.Request("GET", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(len(r.BodyJson().([]any)), 1)
	})

	t.Run("Create Post - Error title validation", func(t *testing.T) {
		r := a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]string{
				"title": strings.Repeat("a", 1025),
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusBadRequest)

		expectedBody := map[string]interface{}{
			"title":       "Bad Request",
			"description": "title is too long",
		}
		biff.AssertEqual(r.BodyJsonMap(), expectedBody)
	})

}
