package algorithms

import (
	"context"
	"fmt"
	"time"

	"sei-ratelimiter/internal/store"
)

type FixedWindow struct {
	store      store.Store
	limit      int
	windowSecs int
}

func NewFixedWindow(store store.Store, limit int, windowSecs int) *FixedWindow {
	return &FixedWindow{
		store:      store,
		limit:      limit,
		windowSecs: windowSecs,
	}
}

func (fw *FixedWindow) Allow(ctx context.Context, clientID string) (bool, int, error) {
	window := time.Now().Unix() / int64(fw.windowSecs)

	key := fmt.Sprintf("fw:%s:%d", clientID, window)

	count, err := fw.store.Incr(ctx, key)
	if err != nil {
		return false, 0, err
	}

	if count == 1 {
		err = fw.store.Expire(ctx, key, fw.windowSecs)
		if err != nil {
			return false, 0, err
		}
	}

	remaining := fw.limit - int(count)

	if int(count) > fw.limit {
		return false, 0, nil
	}

	return true, remaining, nil
}
