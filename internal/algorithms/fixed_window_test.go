package algorithms

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFixedWindow_ConcurrentAllows(t *testing.T) {
	limit := 100
	fw := NewFixedWindow(limit, 60)

	var wg sync.WaitGroup
	var allowed int64

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok, _, _ := fw.Allow(context.Background(), "client-concurrent")
			if ok {
				atomic.AddInt64(&allowed, 1)
			}
		}()
	}

	wg.Wait()

	if int(allowed) != limit {
		t.Errorf("expected exactly %d allowed, got %d", limit, allowed)
	}
}
