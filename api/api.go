package api

import (
	"context"
	"net/http"

	"github.com/fulldump/box"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/statics"
)

func NewApi(version, staticsDir string, p persistence.Persistencer[Post]) http.Handler {

	b := box.NewBox()

	b.WithInterceptors(
		InterceptorAccessLog,
		PrettyError,
	)

	b.WithInterceptors(func(next box.H) box.H {
		return func(ctx context.Context) {
			ctx = context.WithValue(ctx, "persistence", p)
			next(ctx)
		}
	})

	b.HandleResourceNotFound = HandleNotFound
	b.HandleMethodNotAllowed = HandleMethodNotAllowed

	b.Handle("GET", "/", HandleRenderHome)
	b.Handle("GET", "/posts/{postId}", HandleRenderPost)
	b.Handle("GET", "/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	}).WithName("version")

	b.Handle("GET", "/sitemap.xml", HandleSitemap)

	v0 := b.Group("/v0").WithInterceptors(auth.Require)
	v0.Handle("GET", "/posts", HandleListPosts)
	v0.Handle("POST", "/posts", HandleCreatePost)
	v0.Handle("GET", "/posts/{postId}", HandleGetPost)
	v0.Handle("PATCH", "/posts/{postId}", HandleModifyPost)
	v0.Handle("DELETE", "/posts/{postId}", HandleDeletePost)

	// openapi
	buildOpenApi(b)

	// Mount statics
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).WithName("serveStatics")

	return b
}
