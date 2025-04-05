package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/statics"
)

func HandleRenderHome(w http.ResponseWriter, r *http.Request) error {

	// Requires recompile to see changes!!!
	b, err := statics.Www.ReadFile("www/index.gohtml")
	if err != nil {
		log.Println(err)
	}

	tmpl, err := template.New("home").Parse(string(b))
	if err != nil {
		log.Println("template 'home':", err)
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Could not render template",
		}
	}

	posts, err := HandleListPosts(w, r)
	if err != nil {
		return err
	}

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
		Posts     []Post
		TagFilter string
	}{
		Posts:     posts,
		TagFilter: tagFilter,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("template 'home':", err)
	}

	return nil
}
