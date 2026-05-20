package workerpool

import (
	"context"
	"sync"
	"testing"
)

// Simple processor: squares integers
func squareProcessor(job Job[int]) Result[int, int] {
	return Result[int, int]{
		Job: job,
		Output: job.Data * job.Data,
		Err:    nil,
	}
}

func TestWorkerPoolBasic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := NewWorkerPool[int, int](3, 10)
	pool.Start(ctx, squareProcessor)

	// Submit 5 jobs
	for i := 1; i <= 5; i++ {
		if err := pool.Submit(Job[int]{ID: i, Data: i}); err != nil {
			t.Fatalf("failed to submit job: %v", err)
		}
	}

	// Collect results in a goroutine
	results := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range pool.Results() {
			if result.Err != nil {
				t.Errorf("job %d failed: %v", result.Job.ID, result.Err)
				continue
			}
			mu.Lock()
			results[result.Job.ID] = result.Output
			mu.Unlock()
		}
	}()

	pool.Stop() // Close jobs, wait for workers, close results
	wg.Wait()   // Wait for result collector

	// Verify all 5 results
	for i := 1; i <= 5; i++ {
		expected := i * i
		if results[i] != expected {
			t.Errorf("job %d: expected %d, got %d", i, expected, results[i])
		}
	}
}

func TestWorkerPoolConcurrentSubmit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := NewWorkerPool[int, int](5, 100)
	pool.Start(ctx, squareProcessor)

	// Submit 100 jobs concurrently from multiple goroutines
	var submitWg sync.WaitGroup
	numJobs := 100
	submitWg.Add(numJobs)

	for i := 0; i < numJobs; i++ {
		go func(id int) {
			defer submitWg.Done()
			pool.Submit(Job[int]{ID: id, Data: id})
		}(i)
	}

	// Wait for all submissions, then stop
	go func() {
		submitWg.Wait()
		pool.Stop()
	}()

	// Count results
	count := 0
	for range pool.Results() {
		count++
	}

	if count != numJobs {
		t.Errorf("expected %d results, got %d", numJobs, count)
	}
}

func TestWorkerPoolSubmitAfterStopReturnsError(t *testing.T) {
	pool := NewWorkerPool[int, int](1, 1)
	pool.Stop()

	if err := pool.Submit(Job[int]{ID: 1, Data: 1}); err == nil {
		t.Fatalf("expected error when submitting to stopped pool")
	}
}

func TestWorkerPoolStopCanBeCalledTwice(t *testing.T) {
	pool := NewWorkerPool[int, int](1, 1)
	pool.Stop()
	pool.Stop()
}
