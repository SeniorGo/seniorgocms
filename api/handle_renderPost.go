package api

import (
	"github.com/SeniorGo/seniorgocms/utils"
	"html/template"
	"log"
	"net/http"

	"github.com/SeniorGo/seniorgocms/statics"
)

func HandleRenderPost(w http.ResponseWriter, r *http.Request) error {
	b, err := statics.Www.ReadFile("www/post/index.gohtml")
	if err != nil {
		log.Println(err)
	}

	tmpl, err := template.New("post").Funcs(template.FuncMap{
		"formatDateES":          formatDateES,
		"convertMarkdownToHTML": utils.ConvertMarkdownToHTML,
	}).Parse(string(b))
	if err != nil {
		log.Println("template 'post':", err)
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Could not render template",
		}
	}

	post, err := HandleGetPost(r.Context(), r)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		log.Println("template 'post':", err)
	}

	return nil
}
