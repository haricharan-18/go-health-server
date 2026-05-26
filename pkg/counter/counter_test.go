package counter

import (
	"sync"
	"testing"
)

func TestSafeCounterConcurrent(t *testing.T) {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	numGoroutines := 100
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * incrementsPerGoroutine
	actual := counter.Value()

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
