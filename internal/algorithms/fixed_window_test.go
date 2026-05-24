package algorithms

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/Zartex-the-art/sei-ratelimiter/internal/store"
)

func TestFixedWindow(t *testing.T) {
	// Check if Redis is reachable before running
	conn, err := net.DialTimeout("tcp", "localhost:6379", 1*time.Second)
	if err != nil {
		t.Skipf("Redis not available at localhost:6379 — skipping integration test: %v", err)
	}
	conn.Close()

	ctx := context.Background()
	redisStore := store.NewRedisStore("localhost:6379")
	limiter := NewFixedWindow(redisStore, 3, 60)

	for i := 1; i <= 5; i++ {
		allowed, remaining, err := limiter.Allow(ctx, "madhu")
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		t.Logf("request=%d allowed=%v remaining=%d", i, allowed, remaining)
	}
}
