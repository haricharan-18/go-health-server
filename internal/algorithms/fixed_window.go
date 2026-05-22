package algorithms

import (
	"context"
	"sync"
	"time"
)

type FixedWindow struct {
	mu          sync.Mutex
	limit       int
	windowSecs  int
	counts      map[string]int
	windowStart time.Time
}

func NewFixedWindow(limit int, windowSecs int) *FixedWindow {
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

	// Reset window if expired
	if time.Since(fw.windowStart) > time.Duration(fw.windowSecs)*time.Second {
		fw.counts = make(map[string]int)
		fw.windowStart = time.Now()
	}

	fw.counts[clientID]++

	if fw.counts[clientID] > fw.limit {
		return false, 0, nil
	}

	remaining := fw.limit - fw.counts[clientID]

	return true, remaining, nil
}