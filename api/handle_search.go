package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	ctx := r.Context()

	p := GetPostPersistence(ctx)

	posts, err := p.List(ctx)
	if err != nil {
		log.Println("p.List:", err)
		return nil, ErrorPersistenceRead
	}

	queryParams := r.URL.Query()
	q := queryParams.Get("q")
	if q == "" {
		http.Error(w, "q is required", http.StatusBadRequest)
		return nil, fmt.Errorf("q is required")
	}

	qLower := strings.ToLower(q)

	var sPosts []Post

	for _, post := range posts {
		titleLower := strings.ToLower(post.Item.Title)
		bodyLower := strings.ToLower(post.Item.Body)
		if strings.Contains(titleLower, qLower) || strings.Contains(bodyLower, qLower) {
			sPosts = append(sPosts, post.Item)
		}
	}

	return sPosts, nil

}
