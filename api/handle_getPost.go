package api

import (
	"context"
)

func HandleGetPost(ctx context.Context) *Post {
	post := GetPost(ctx) // get post from context
	return &post.Item
}
