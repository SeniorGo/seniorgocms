package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/biff"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
)

func TestNewApi(t *testing.T) {

	// postPersistencer := persistence.NewInMemory[Post]()
	postPersistencer, err := persistence.NewInDisk[Post](t.TempDir())
	biff.AssertNil(err)

	categoryPersistencer, err := persistence.NewInDisk[Category](t.TempDir())
	biff.AssertNil(err)

	user := auth.User{
		ID:      "user-test",
		Nick:    "user-nick",
		Picture: "user-picture",
		Email:   "user@email.com",
	}

	authHeaderBytes, _ := json.Marshal(map[string]any{"user": user})
	authHeader := string(authHeaderBytes)

	// Create a test logger
	testLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	h := NewApi("testversion", "", postPersistencer, categoryPersistencer, testLogger)
	a := apitest.NewWithHandler(h)

	t.Run("Request /version", func(t *testing.T) {
		r := a.Request("GET", "/version").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(r.BodyString(), "testversion")
	})

	t.Run("List posts (empty)", func(t *testing.T) {
		r := a.Request("GET", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)

		//expectedBody := []map[string]interface{}{}
		expectedBody := map[string]interface{}{
			"total": 0,
			"posts": []map[string]interface{}{},
			"limit": 10,
			"skip":  0,
		}
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
			"author":            user,
			"title":             "My Post",
			"body":              "This is my body",
			"tags":              nil,
			"creation_time":     body["creation_time"],
			"modification_time": body["modification_time"],
		}
		biff.AssertEqualJson(r.BodyJsonMap(), expectedBody)
	})

	t.Run("List posts (1)", func(t *testing.T) {
		r := a.Request("GET", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			Do()

		posts := r.BodyJson().(map[string]interface{})["posts"].([]interface{})

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertEqual(len(posts), 1)
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
			"error": map[string]interface{}{
				"title":       "Bad Request",
				"description": "title is too long",
			},
		}
		biff.AssertEqual(r.BodyJsonMap(), expectedBody)
	})

	t.Run("Create Category", func(t *testing.T) {
		r := a.Request("POST", "/v0/categories").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]string{
				"name": "My Category",
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusCreated)

		body := r.BodyJsonMap()
		expectedBody := map[string]interface{}{
			"id":                body["id"],
			"name":              "My Category",
			"creation_time":     body["creation_time"],
			"modification_time": body["modification_time"],
		}
		biff.AssertEqualJson(r.BodyJsonMap(), expectedBody)
	})

	t.Run("Create Post with Tags", func(t *testing.T) {
		r := a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Tagged Post",
				"body":  "This post has tags",
				"tags":  []string{"test", "golang"},
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusCreated)

		body := r.BodyJsonMap()
		expectedBody := map[string]interface{}{
			"id":                body["id"],
			"author":            user,
			"title":             "Tagged Post",
			"body":              "This post has tags",
			"tags":              []string{"test", "golang"},
			"creation_time":     body["creation_time"],
			"modification_time": body["modification_time"],
		}
		biff.AssertEqualJson(r.BodyJsonMap(), expectedBody)
	})

	t.Run("Create Post - Error tag validation", func(t *testing.T) {
		r := a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Post with long tag",
				"body":  "This post has a too long tag",
				"tags":  []string{strings.Repeat("a", 129)},
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusBadRequest)

		expectedBody := map[string]interface{}{
			"error": map[string]interface{}{
				"title":       "Bad Request",
				"description": "tag is too long (max 128 characters)",
			},
		}
		biff.AssertEqual(r.BodyJsonMap(), expectedBody)
	})

	t.Run("Modify Post Tags", func(t *testing.T) {
		// First create a post
		createResp := a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Post to modify tags",
				"body":  "This post will have its tags modified",
				"tags":  []string{"initial"},
			}).
			Do()

		postId := createResp.BodyJsonMap()["id"].(string)

		// Now modify the tags
		r := a.Request("PATCH", "/v0/posts/"+postId).
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"tags": []string{"updated", "tags"},
			}).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)

		body := r.BodyJsonMap()
		biff.AssertEqual(body["tags"], []interface{}{"updated", "tags"})
	})

	t.Run("Search Posts By Query", func(t *testing.T) {
		a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Post to search posts",
				"body":  "This post with any body only for search with query param",
			}).
			Do()

		query := "param"
		r := a.Request("GET", "/v0/search").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithQuery("q", query).
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)

		posts := r.BodyJson().([]any)
		biff.AssertTrue(len(posts) > 0)

		found := false
		for _, p := range posts {
			postMap := p.(map[string]interface{})
			title := postMap["title"].(string)
			body := postMap["body"].(string)
			if strings.Contains(strings.ToLower(title), query) || strings.Contains(strings.ToLower(body), query) {
				found = true
				break
			}
		}
		biff.AssertTrue(found)
	})

	t.Run("List Posts with Tag Filter", func(t *testing.T) {
		// Create posts with different tags
		a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Post with tag A",
				"body":  "Content A",
				"tags":  []string{"tagA", "common"},
			}).
			Do()

		a.Request("POST", "/v0/posts").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithBodyJson(map[string]any{
				"title": "Post with tag B",
				"body":  "Content B",
				"tags":  []string{"tagB", "common"},
			}).
			Do()

		// Test filtering by tagA
		r := a.Request("GET", "/").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithQuery("tag", "tagA").
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		// Response should contain the HTML with filtered posts
		biff.AssertTrue(strings.Contains(r.BodyString(), "Post with tag A"))
		biff.AssertFalse(strings.Contains(r.BodyString(), "Post with tag B"))

		// Test filtering by common tag
		r = a.Request("GET", "/").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithQuery("tag", "common").
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		// Response should contain both posts
		biff.AssertTrue(strings.Contains(r.BodyString(), "Post with tag A"))
		biff.AssertTrue(strings.Contains(r.BodyString(), "Post with tag B"))

		// Test filtering by non-existent tag
		r = a.Request("GET", "/").
			WithHeader(auth.XGlueAuthentication, authHeader).
			WithQuery("tag", "nonexistent").
			Do()

		biff.AssertEqual(r.StatusCode, http.StatusOK)
		biff.AssertTrue(strings.Contains(r.BodyString(), "No hay posts con la etiqueta"))
	})
}
