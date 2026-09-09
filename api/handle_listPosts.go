package api

import (
	"log"
	"net/http"
	"sort"
	"strconv"
)

type PaginatedPosts struct {
	Total int    `json:"total"`
	Limit int    `json:"limit"`
	Skip  int    `json:"skip"`
	Posts []Post `json:"posts"`
}

func HandleListPosts(w http.ResponseWriter, r *http.Request) (PaginatedPosts, error) {

	ctx := r.Context()

	p := GetPostPersistence(ctx)

	posts, err := p.List(ctx)
	if err != nil {
		log.Println("p.List:", err)
		return PaginatedPosts{}, ErrorPersistenceRead
	}

	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = post.Item
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreationTime.After(result[j].CreationTime)
	})

	queryParams := r.URL.Query()
	limitStr := queryParams.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	skipStr := queryParams.Get("skip")
	skip, err := strconv.Atoi(skipStr)
	if err != nil || skip < 0 {
		skip = 0
	}

	total := len(result)

	start := skip
	if start > len(result) {
		start = len(result)
	}

	end := start + limit
	if end > len(result) {
		end = len(result)
	}

	result = result[start:end]

	return PaginatedPosts{
		Total: total,
		Limit: limit,
		Skip:  skip,
		Posts: result,
	}, nil
}
