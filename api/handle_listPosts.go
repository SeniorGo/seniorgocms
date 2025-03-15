package api

import (
	"log"
	"net/http"
	"sort"

	"github.com/SeniorGo/seniorgocms/persistence"
)

type listPosts struct {
	postRepo persistence.Persistencer[Post]
}

func newListPosts(postRepo persistence.Persistencer[Post]) *listPosts {
	return &listPosts{postRepo: postRepo}
}

func (h *listPosts) Handle(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	ctx := r.Context()

	posts, err := h.postRepo.List(ctx)
	if err != nil {
		log.Println("h.postRepo.List:", err)
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
