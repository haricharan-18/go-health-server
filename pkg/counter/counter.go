package counter

import "sync"

// SafeCounter is a thread-safe counter using sync.Mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// NewSafeCounter creates a new SafeCounter
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Inc increments the counter safely
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current counter value safely
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// UnsafeCounter demonstrates what NOT to do — no mutex protection
type UnsafeCounter struct {
	value int
}

func (c *UnsafeCounter) Inc() {
	c.value++
}

func (c *UnsafeCounter) Value() int {
	return c.value
}