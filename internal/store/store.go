package store

import "context"

type Store interface {
	Increment(ctx context.Context, key string, windowSecs int) (int64, error)

	ZAdd(ctx context.Context, key string, score float64, member string) error
	ZRemRangeByScore(ctx context.Context, key string, min, max float64) error
	ZCount(ctx context.Context, key string, min, max float64) (int64, error)

	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HSet(ctx context.Context, key string, values map[string]interface{}) error

	Del(ctx context.Context, keys ...string) error
}
