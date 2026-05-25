package algorithms

import "context"

// Limiter interface for all rate limiting algorithms
type Limiter interface {
	Allow(ctx context.Context, clientID string) (bool, int, error)
         Reset(clientID string)
}
