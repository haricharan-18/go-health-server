package workerpool

import (
	"context"
	"fmt"
	"sync"
)

// Job represents a unit of work
type Job[T any] struct {
	ID   int
	Data T
}

// Result represents the output of processing a Job
type Result[T any, R any] struct {
	Job    Job[T]
	Output R
	Err    error
}

// WorkerPool manages a pool of goroutines that process jobs
type WorkerPool[T any, R any] struct {
	numWorkers int
	jobs       chan Job[T]
	results    chan Result[T, R]
	wg         sync.WaitGroup
	mu         sync.Mutex
	closed     bool
	stopOnce   sync.Once
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool[T any, R any](numWorkers, jobQueueSize int) (*WorkerPool[T, R], error) {
	if numWorkers <= 0 {
		return nil, fmt.Errorf("numWorkers must be > 0")
	}
	if jobQueueSize < 0 {
		return nil, fmt.Errorf("jobQueueSize must be >= 0")
	}

	return &WorkerPool[T, R]{
		numWorkers: numWorkers,
		jobs:       make(chan Job[T], jobQueueSize),
		results:    make(chan Result[T, R], jobQueueSize),
	}, nil
}

// Start initializes the worker goroutines
func (wp *WorkerPool[T, R]) Start(ctx context.Context, processor func(Job[T]) Result[T, R]) {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case job, ok := <-wp.jobs:
					if !ok {
						return
					}
					result := processor(job)
					select {
					case wp.results <- result:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

// Submit adds a job to the queue
func (wp *WorkerPool[T, R]) Submit(job Job[T]) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.closed {
		return fmt.Errorf("worker pool is closed")
	}

	select {
	case wp.jobs <- job:
		return nil
	default:
		return fmt.Errorf("job queue full")
	}
}

// Results returns the results channel for reading
func (wp *WorkerPool[T, R]) Results() <-chan Result[T, R] {
	return wp.results
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool[T, R]) Stop() {
	wp.stopOnce.Do(func() {
		wp.mu.Lock()
		wp.closed = true
		close(wp.jobs)
		wp.mu.Unlock()

		wp.wg.Wait()
		close(wp.results)
	})
}
