package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/statics"
)

type renderHome struct {
	postRepo persistence.Persistencer[Post]
}

func newRenderHome(postRepo persistence.Persistencer[Post]) *renderHome {
	return &renderHome{postRepo: postRepo}
}

func (h *renderHome) Handle(w http.ResponseWriter, r *http.Request) error {
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

	posts, err := newListPosts(h.postRepo).Handle(w, r)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		log.Println("template 'home':", err)
	}

	return nil
}
