package persistence

import (
	"context"
	"sync"
)

type InMemory[T any] struct {
	Items map[string]*ItemWithId[T]
	mutex sync.RWMutex
}

func NewInMemory[T any]() *InMemory[T] {
	return &InMemory[T]{
		Items: map[string]*ItemWithId[T]{},
	}
}

func (f *InMemory[T]) List(ctx context.Context) ([]*ItemWithId[T], error) {

	f.mutex.RLock()
	defer f.mutex.RUnlock()

	result := make([]*ItemWithId[T], len(f.Items))

	i := -1
	for _, f := range f.Items {
		i++
		result[i] = f
	}

	return result, nil
}

func (f *InMemory[T]) Put(ctx context.Context, item *ItemWithId[T]) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if oldItem, ok := f.Items[item.Id]; ok {
		if oldItem.Version != item.Version {
			return ErrVersionGone
		}
	}

	f.Items[item.Id] = &ItemWithId[T]{
		Id:      item.Id,
		Item:    item.Item,
		Version: item.Version + 1,
	}
	return nil
}

func (f *InMemory[T]) Get(ctx context.Context, id string) (*ItemWithId[T], error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	item, ok := f.Items[id]
	if !ok {
		return nil, nil
	}

	return &ItemWithId[T]{
		Id:      item.Id,
		Item:    item.Item,
		Version: item.Version,
	}, nil
}

func (f *InMemory[T]) Delete(ctx context.Context, id string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	delete(f.Items, id)
	return nil
}
