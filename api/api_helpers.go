package api

import (
	"context"

	"github.com/SeniorGo/seniorgocms/persistence"
)

// GetPersistence from context (where posts are stored)
func GetPersistence(ctx context.Context) persistence.Persistencer[Post] {
	p, ok := ctx.Value("persistence").(persistence.Persistencer[Post])
	if !ok {
		panic("Persistence should be in context")
	}

	return p
}
