package workerpool

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

// Simple processor: squares integers
func squareProcessor(job Job) Result {
	if num, ok := job.Data.(int); ok {
		return Result{
			Job:    job,
			Output: num * num,
			Err:    nil,
		}
	}
	return Result{
		Job: job,
		Err: fmt.Errorf("invalid data type"),
	}
}

func TestWorkerPoolBasic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := NewWorkerPool(3, 10)
	pool.Start(ctx, squareProcessor)

	// Submit 5 jobs
	for i := 1; i <= 5; i++ {
		if err := pool.Submit(Job{ID: i, Data: i}); err != nil {
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
			results[result.Job.ID] = result.Output.(int)
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

	pool := NewWorkerPool(5, 100)
	pool.Start(ctx, squareProcessor)

	// Submit 100 jobs concurrently from multiple goroutines
	var submitWg sync.WaitGroup
	numJobs := 100
	submitWg.Add(numJobs)

	for i := 0; i < numJobs; i++ {
		go func(id int) {
			defer submitWg.Done()
			pool.Submit(Job{ID: id, Data: id})
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