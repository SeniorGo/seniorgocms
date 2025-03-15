package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/statics"
)

type renderPost struct {
	postRepo persistence.Persistencer[Post]
}

func newRenderPost(postRepo persistence.Persistencer[Post]) *renderPost {
	return &renderPost{postRepo: postRepo}
}

func (h *renderPost) Handle(w http.ResponseWriter, r *http.Request) error {
	b, err := statics.Www.ReadFile("www/post/index.gohtml")
	if err != nil {
		log.Println(err)
	}

	tmpl, err := template.New("post").Funcs(template.FuncMap{
		"formatDateES": formatDateES,
	}).Parse(string(b))
	if err != nil {
		log.Println("template 'post':", err)
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Could not render template",
		}
	}

	post, err := newGetPost(h.postRepo).Handle(r.Context(), r)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		log.Println("template 'post':", err)
	}

	return nil
}
