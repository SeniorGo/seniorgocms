package api

import (
	"context"
	"github.com/SeniorGo/seniorgocms/logger"
	"log/slog"
	"net/http"

	"github.com/fulldump/box"

	"github.com/SeniorGo/seniorgocms/auth"
	"github.com/SeniorGo/seniorgocms/persistence"
	"github.com/SeniorGo/seniorgocms/statics"
)

func NewApi(
	version, staticsDir string,
	postPersistencer persistence.Persistencer[Post],
	categoryPersistencer persistence.Persistencer[Category],
	log *slog.Logger,
) http.Handler {

	b := box.NewBox()

	b.WithInterceptors(
		logger.InjectLog(log),
		InterceptorAccessLog,
		PrettyError,
	)

	b.WithInterceptors(func(next box.H) box.H {
		return func(ctx context.Context) {
			ctx = context.WithValue(ctx, "post-persistence", postPersistencer)
			ctx = context.WithValue(ctx, "category-persistence", categoryPersistencer)
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
	v0.Handle("POST", "/categories", HandleCreateCategory)

	// openapi
	buildOpenApi(b)

	// Mount statics
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).WithName("serveStatics")

	return b
}
