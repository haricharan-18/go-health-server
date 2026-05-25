package algorithms

import (
	"context"
	"sync"
	"time"
)

type FixedWindow struct {
	limit      int
	windowSecs int
	mu         sync.Mutex
	counts     map[string]int
	windows    map[string]time.Time
}

func NewFixedWindow(_ interface{}, limit, windowSecs int) *FixedWindow {
	return &FixedWindow{
		limit:      limit,
		windowSecs: windowSecs,
		counts:     make(map[string]int),
		windows:    make(map[string]time.Time),
	}
}

func (fw *FixedWindow) Allow(ctx context.Context, clientID string) (bool, int, error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()

	windowStart, exists := fw.windows[clientID]

	if !exists || now.Sub(windowStart) >= time.Duration(fw.windowSecs)*time.Second {
		fw.counts[clientID] = 0
		fw.windows[clientID] = now
	}

	fw.counts[clientID]++

	count := fw.counts[clientID]

	remaining := fw.limit - count

	if remaining < 0 {
		remaining = 0
	}

	if count > fw.limit {
		return false, 0, nil
	}

	return true, remaining, nil
}

func (fw *FixedWindow) Reset(clientID string) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	delete(fw.counts, clientID)
	delete(fw.windows, clientID)
}
