package store

import (
	"context"
	"sync"
)

type FakeStore struct {
	mu   sync.Mutex
	data map[string]int64
}

func NewFakeStore() *FakeStore {
	return &FakeStore{
		data: make(map[string]int64),
	}
}

func (f *FakeStore) Increment(ctx context.Context, key string, windowSecs int) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.data[key]++
	return f.data[key], nil
}

func (f *FakeStore) Reset(key string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	delete(f.data, key)
}

func (f *FakeStore) Del(ctx context.Context, keys ...string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, k := range keys {
		delete(f.data, k)
	}
	return nil
}

func (f *FakeStore) ZAdd(ctx context.Context, key string, score float64, member string) error {
	return nil
}

func (f *FakeStore) ZRemRangeByScore(ctx context.Context, key string, min, max float64) error {
	return nil
}

func (f *FakeStore) ZCount(ctx context.Context, key string, min, max float64) (int64, error) {
	return 0, nil
}

func (f *FakeStore) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return map[string]string{}, nil
}

func (f *FakeStore) HSet(ctx context.Context, key string, values map[string]interface{}) error {
	return nil
}
