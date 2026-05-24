package store

import "context"

type Store interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, seconds int) error
}
