package algorithms

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"sei-ratelimiter/internal/store"
)

// TestFixedWindow_Integration validates basic Redis-backed behavior
func TestFixedWindow_Integration(t *testing.T) {
	ctx := context.Background()

	redisStore := store.NewRedisStore("localhost:6379")
	limiter := NewFixedWindow(redisStore, 3, 60)

	for i := 1; i <= 5; i++ {
		allowed, remaining, err := limiter.Allow(ctx, "madhu")
		if err != nil {
			t.Fatalf("request %d: error: %v", i, err)
		}

		t.Logf("request=%d allowed=%v remaining=%d", i, allowed, remaining)

		// Assertions for clarity
		if i <= 3 && !allowed {
			t.Errorf("request %d: expected allowed=true within limit", i)
		}
		if i > 3 && allowed {
			t.Errorf("request %d: expected allowed=false over limit", i)
		}
	}
}

// TestFixedWindow_Concurrent stress-tests thread safety
func TestFixedWindow_Concurrent(t *testing.T) {
	ctx := context.Background()

	// Use in-memory store for deterministic concurrent testing
	memStore := store.NewMemoryStore()
	limit := 100
	limiter := NewFixedWindow(memStore, limit, 60)

	var wg sync.WaitGroup
	var allowed int64

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok, _, _ := limiter.Allow(ctx, "client-1")
			if ok {
				atomic.AddInt64(&allowed, 1)
			}
		}()
	}

	wg.Wait()

	if int(allowed) != limit {
		t.Errorf("expected %d allowed, got %d", limit, allowed)
	}
}