package api

import (
	"log"
	"net/http"
	"sort"
)

func HandleListPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {

	ctx := r.Context()

	p := GetPostPersistence(ctx)

	posts, err := p.List(ctx)
	if err != nil {
		log.Println("p.List:", err)
		return nil, ErrorPersistenceRead
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
