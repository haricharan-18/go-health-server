package algorithms

import (
	"context"
	"fmt"

	"github.com/Zartex-the-art/sei-ratelimiter/internal/store"
)

type FixedWindow struct {
	limit      int
	windowSecs int
	store      store.Store
}

func NewFixedWindow(limit, windowSecs int, s store.Store) *FixedWindow {
	return &FixedWindow{
		limit:      limit,
		windowSecs: windowSecs,
		store:      s,
	}
}
func (fw *FixedWindow) Allow(ctx context.Context, clientID string) (bool, int, error) {
	key := fmt.Sprintf("fw:%s", clientID)

	count, err := fw.store.Increment(
		ctx,
		key,
		fw.windowSecs,
	)
	if err != nil {
		return false, 0, err
	}

	remaining := fw.limit - int(count)

	if remaining < 0 {
		remaining = 0
	}

	if int(count) > fw.limit {
		return false, 0, nil
	}

	return true, remaining, nil
}
