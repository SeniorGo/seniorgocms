package persistence

import (
	"context"
	"errors"
)

type ItemWithId[T any] struct {
	Id      string `json:"id"`
	Item    T      `bson:",inline"`
	Version int64  `json:"-" bson:"version"`
}

var ErrVersionGone = errors.New("version gone")

type Persistencer[T any] interface {
	List(ctx context.Context) ([]*ItemWithId[T], error)
	Put(ctx context.Context, item *ItemWithId[T]) error
	Get(ctx context.Context, id string) (*ItemWithId[T], error)
	Delete(ctx context.Context, id string) error
}
