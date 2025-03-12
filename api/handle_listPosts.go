package api

import (
	"log"
	"net/http"
	"sort"
)

func HandleListPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {

	ctx := r.Context()

	p := GetPersistence(ctx)

	posts, err := p.List(ctx)
	if err != nil {
		log.Println("p.List:", err)
		return nil, HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Problem reading from persistence layer",
		}
	}

	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = post.Item
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreationTime.After(result[j].CreationTime)
	})

	return result, nil
}
