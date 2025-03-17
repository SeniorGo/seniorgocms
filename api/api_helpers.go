package api

import (
	"context"

	"github.com/SeniorGo/seniorgocms/persistence"
)

// GetPostPersistence from context (where posts are stored)
func GetPostPersistence(ctx context.Context) persistence.Persistencer[Post] {
	p, ok := ctx.Value("post-persistence").(persistence.Persistencer[Post])
	if !ok {
		panic("Persistence should be in context")
	}

	return p
}

// GetCategoryPersistence from context (where categories are stored)
func GetCategoryPersistence(ctx context.Context) persistence.Persistencer[Category] {
	p, ok := ctx.Value("category-persistence").(persistence.Persistencer[Category])
	if !ok {
		panic("Persistence should be in context")
	}

	return p
}
