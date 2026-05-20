package algorithms

import (
	"context"
	"sync"
	"time"
)

type FixedWindow struct {
	limit       int
	windowSecs  int
	mu          sync.Mutex
	counts      map[string]int
	windowStart time.Time
}

func NewFixedWindow(limit, windowSecs int) *FixedWindow {
	return &FixedWindow{
		limit:       limit,
		windowSecs:  windowSecs,
		counts:      make(map[string]int),
		windowStart: time.Now(),
	}
}

func (fw *FixedWindow) Allow(ctx context.Context, clientID string) (bool, int, error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if time.Since(fw.windowStart) > time.Duration(fw.windowSecs)*time.Second {
		fw.counts = make(map[string]int)
		fw.windowStart = time.Now()
	}

	if fw.counts[clientID] >= fw.limit {
		return false, 0, nil
	}

	fw.counts[clientID]++

	remaining := fw.limit - fw.counts[clientID]

	return true, remaining, nil
}
package algorithms

import (
    "context"
    "sync"
    "sync/atomic"
    "testing"
)func TestFixedWindow_ConcurrentAllows(t *testing.T) {
    limit := 100
    fw := NewFixedWindow(limit, 60)

    var wg sync.WaitGroup
    var allowed int64  // use atomic int for the test counter itself

    // Launch 300 goroutines — 3x the limit
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

    // Exactly 'limit' requests must have been allowed — not more, not less
    if int(allowed) != limit {
        t.Errorf("expected exactly %d allowed, got %d", limit, allowed)
    }
}