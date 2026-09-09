package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/statics"
	"strconv"
)

func HandleRenderHome(w http.ResponseWriter, r *http.Request) error {

	// Requires recompile to see changes!!!
	b, err := statics.Www.ReadFile("www/index.gohtml")
	if err != nil {
		log.Println(err)
	}

	tmpl, err := template.New("home").Funcs(template.FuncMap{
		"toInt": func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		},
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"until": func(n int) []int {
			arr := make([]int, n)
			for i := 0; i < n; i++ {
				arr[i] = i
			}
			return arr
		},
	}).Parse(string(b))
	if err != nil {
		log.Println("template 'home':", err)
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Could not render template",
		}
	}

	paginatedPosts, err := HandleListPosts(w, r)
	if err != nil {
		return err
	}

	posts := paginatedPosts.Posts
	// Filter posts by tag if tag parameter is present
	tagFilter := r.URL.Query().Get("tag")
	if tagFilter != "" {
		filteredPosts := make([]Post, 0)
		for _, post := range posts {
			for _, tag := range post.Tags {
				if tag == tagFilter {
					filteredPosts = append(filteredPosts, post)
					break
				}
			}
		}
		posts = filteredPosts
	}

	data := struct {
		Posts      []Post
		TotalPosts int
		Limit      int
		Skip       int
		TagFilter  string
	}{
		Posts:      posts,
		TotalPosts: paginatedPosts.Total,
		Limit:      paginatedPosts.Limit,
		Skip:       paginatedPosts.Skip,
		TagFilter:  tagFilter,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("template 'home':", err)
	}

	return nil
}
