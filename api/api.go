package api

import (
	"net/http"

	"github.com/fulldump/box"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/statics"
)

func NewApi(version, staticsDir string, postRepo persistence.Persistencer[Post]) http.Handler {

	b := box.NewBox()

	b.WithInterceptors(
		InterceptorAccessLog,
		PrettyError,
	)

	b.HandleResourceNotFound = HandleNotFound
	b.HandleMethodNotAllowed = HandleMethodNotAllowed

	b.Handle("GET", "/", newRenderHome(postRepo).Handle)
	b.Handle("GET", "/posts/{postId}", newRenderPost(postRepo).Handle)
	b.Handle("GET", "/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	}).WithName("version")

	b.Handle("GET", "/sitemap.xml", newSitemap(postRepo).Handle)

	v0 := b.Group("/v0").WithInterceptors(auth.Require)
	v0.Handle("GET", "/posts", newListPosts(postRepo).Handle)
	v0.Handle("POST", "/posts", newCreatePost(postRepo).Handle)
	v0.Handle("GET", "/posts/{postId}", newGetPost(postRepo).Handle)
	v0.Handle("PATCH", "/posts/{postId}", newModifyPost(postRepo).Handle)
	v0.Handle("DELETE", "/posts/{postId}", newDeletePost(postRepo).Handle)

	// openapi
	buildOpenApi(b)

	// Mount statics
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).WithName("serveStatics")

	return b
}
