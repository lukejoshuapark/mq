package complex

import (
	"context"
	"time"
)

type QueryOptions struct {
	Limit  int
	Offset int
	Sort   string
}

type Result[T any] struct {
	Data  T
	Error error
}

type DataStore interface {
	Query(ctx context.Context, query string, options *QueryOptions) (map[string]interface{}, error)
	Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	GetMultiple(ctx context.Context, keys []string) ([]interface{}, error)
	ProcessBatch(ctx context.Context, items ...interface{}) (int, error)
}
