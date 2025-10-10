package genericinterface

import "context"

type Repository[T any] interface {
	Get(ctx context.Context, id string) (T, error)
	Save(ctx context.Context, item T) error
	List(ctx context.Context) ([]T, error)
}

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K) bool
}
