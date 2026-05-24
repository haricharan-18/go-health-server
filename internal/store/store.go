package store

import "context"

type Store interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, seconds int) error
}

// MemoryStore is a simple in-memory implementation for testing.
type MemoryStore struct {
	data map[string]int64
}

// NewMemoryStore creates a new in-memory store for testing.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]int64)}
}

func (m *MemoryStore) Incr(ctx context.Context, key string) (int64, error) {
	m.data[key]++
	return m.data[key], nil
}

func (m *MemoryStore) Expire(ctx context.Context, key string, seconds int) error {
	return nil
}
