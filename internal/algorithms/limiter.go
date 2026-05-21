package algorithms

import "context"

// Limiter interface
type Limiter interface {
	Allow(ctx context.Context, clientID string) (bool, int, error)
}

// Config for limiter
type Config struct {
	Algorithm  string
	Limit      int
	WindowSecs int
}
