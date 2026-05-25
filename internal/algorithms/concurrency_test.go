package algorithms

import (
	"context"
	"testing"
)

type fakeStore struct {
	count int64
}

func (f *fakeStore) Incr(ctx context.Context, key string) (int64, error) {
	f.count++
	return f.count, nil
}

func (f *fakeStore) Expire(ctx context.Context, key string, seconds int) error {
	return nil
}

func TestFixedWindow_AllowsUnderLimit(t *testing.T) {
	store := &fakeStore{}
	fw := NewFixedWindow(store, 5, 60)
	
	ctx := context.Background()
	allowed, remaining, err := fw.Allow(ctx, "client-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !allowed {
		t.Error("expected allowed=true for first request")
	}
	if remaining != 4 {
		t.Errorf("expected remaining=4, got %d", remaining)
	}
}

func TestFixedWindow_BlocksAtLimit(t *testing.T) {
	store := &fakeStore{}
	fw := NewFixedWindow(store, 2, 60)
	ctx := context.Background()
	
	// Make 3 requests, only 2 should pass
	fw.Allow(ctx, "client-1")
	fw.Allow(ctx, "client-1")
	allowed, _, err := fw.Allow(ctx, "client-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if allowed {
		t.Error("expected allowed=false when limit exceeded")
	}
}
