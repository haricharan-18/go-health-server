package algorithms

import (
	"context"
	"testing"

	"sei-ratelimiter/internal/store"
)

func TestFixedWindow(t *testing.T) {
	ctx := context.Background()

	redisStore := store.NewRedisStore("localhost:6379")

	limiter := NewFixedWindow(redisStore, 3, 60)

	for i := 1; i <= 5; i++ {
		allowed, remaining, err := limiter.Allow(ctx, "madhu")

		if err != nil {
			t.Fatalf("error: %v", err)
		}

		t.Logf(
			"request=%d allowed=%v remaining=%d",
			i,
			allowed,
			remaining,
		)
	}
}
