package algorithms

import "context"

type Limiter interface {
	Allow(ctx context.Context, clientID string) (bool, int, error)
}