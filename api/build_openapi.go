package api

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/box"
	"github.com/fulldump/box/boxopenapi"
)

func buildOpenApi(b *box.B) {
	spec := boxopenapi.Spec(b)
	spec.Info.Title = "SeniorGoCMS"
	spec.Info.Description = "A free CMS for learning and more!"
	spec.Info.Contact = &boxopenapi.Contact{
		Url: "https://github.com/SeniorGo/seniorgocms/issues/new",
	}
	b.Handle("GET", "/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		spec.Servers = []boxopenapi.Server{
			{
				Url: "https://" + r.Host,
			},
			{
				Url: "http://" + r.Host,
			},
		}

		e := json.NewEncoder(w)
		e.SetIndent("", "    ")
		e.Encode(spec)
	}).WithName("OpenApi")
}
