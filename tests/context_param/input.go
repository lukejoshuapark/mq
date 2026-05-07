package contextparam

import "context"

type Service interface {
	Fetch(ctx context.Context, id string) (string, error)
	Store(ctx context.Context, key string, value string) error
	Ping(ctx context.Context) error
}
