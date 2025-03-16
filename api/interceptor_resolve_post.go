package api

import (
	"context"
	"log"

	"github.com/fulldump/box"
)

// ResolvePost fetches the post from the database looking for 'postId' path value.
// Also validates user input and handle errors (database and not found)
func ResolvePost(next box.H) box.H {
	return func(ctx context.Context) {
		r := box.GetRequest(ctx)

		postId := r.PathValue("postId")

		post, err := GetPersistence(ctx).Get(ctx, postId)
		if err != nil {
			log.Println("p.Get:", err)
			box.SetError(ctx, ErrorPersistenceRead)
			return
		}
		if post == nil {
			box.SetError(ctx, ErrorPostNotFound)
			return
		}

		ctx = context.WithValue(ctx, "post", post)
		next(ctx)
	}
}
