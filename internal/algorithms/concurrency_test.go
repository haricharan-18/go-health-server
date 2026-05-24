package algorithms

import (
    "context"
    "sync"
    "sync/atomic"
    "testing"
)

// TestConcurrentRequests_AllowedCountIsExact verifies that under concurrent load,
// the rate limiter allows exactly 'limit' requests and no more.
func TestConcurrentRequests_AllowedCountIsExact(t *testing.T) {
    const limit = 50
    const goroutines = 200

    fw := NewFixedWindow(limit, 60)

    var wg sync.WaitGroup
    var allowed, blocked int64

    for i := 0; i < goroutines; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            ok, _, err := fw.Allow(context.Background(), "shared-client")

            if err != nil {
                t.Errorf("unexpected error: %v", err)
                return
            }

            if ok {
                atomic.AddInt64(&allowed, 1)
            } else {
                atomic.AddInt64(&blocked, 1)
            }
        }()
    }

    wg.Wait()

    if int(allowed) != limit {
        t.Errorf("allowed: got %d, want %d", allowed, limit)
    }

    if int(blocked) != goroutines-limit {
        t.Errorf("blocked: got %d, want %d", blocked, goroutines-limit)
    }

    if allowed+blocked != int64(goroutines) {
        t.Errorf("allowed + blocked = %d, want %d", allowed+blocked, goroutines)
    }
}

// TestConcurrentRequests_MultipleClients verifies independent limits per client
func TestConcurrentRequests_MultipleClients(t *testing.T) {
    const limit = 20

    fw := NewFixedWindow(limit, 60)

    var wg sync.WaitGroup
    var allowed1, allowed2 int64

    for i := 0; i < 60; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            ok, _, _ := fw.Allow(context.Background(), "client-A")

            if ok {
                atomic.AddInt64(&allowed1, 1)
            }
        }()
    }

    for i := 0; i < 60; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            ok, _, _ := fw.Allow(context.Background(), "client-B")

            if ok {
                atomic.AddInt64(&allowed2, 1)
            }
        }()
    }

    wg.Wait()

    if int(allowed1) != limit {
        t.Errorf("client-A: got %d, want %d", allowed1, limit)
    }

    if int(allowed2) != limit {
        t.Errorf("client-B: got %d, want %d", allowed2, limit)
    }
}