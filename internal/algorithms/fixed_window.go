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
