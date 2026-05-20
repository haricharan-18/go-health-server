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
	done       chan struct{}
	wg         sync.WaitGroup
	mu         sync.RWMutex
	closed     bool
	stopOnce   sync.Once
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool[T any, R any](numWorkers, jobQueueSize int) *WorkerPool[T, R] {
	if numWorkers <= 0 {
		numWorkers = 1
	}
	if jobQueueSize <= 0 {
		jobQueueSize = 1
	}

	return &WorkerPool[T, R]{
		numWorkers: numWorkers,
		jobs:       make(chan Job[T], jobQueueSize),
		results:    make(chan Result[T, R], jobQueueSize),
		done:       make(chan struct{}),
	}
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
						return // Channel closed, worker exits
					}
					result := processor(job)
					select {
					case wp.results <- result:
					case <-ctx.Done():
						return
					default:
						select {
						case wp.results <- result:
						case <-ctx.Done():
							return
						case <-wp.done:
							return
						}
					}
				case <-ctx.Done():
					return // Context cancelled, worker exits
				}
			}
		}()
	}
}

// Submit adds a job to the queue
func (wp *WorkerPool[T, R]) Submit(job Job[T]) error {
	wp.mu.RLock()
	defer wp.mu.RUnlock()

	if wp.closed {
		return fmt.Errorf("worker pool stopped")
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
		close(wp.jobs) // No more jobs accepted
		close(wp.done)
		wp.mu.Unlock()

		wp.wg.Wait()      // Wait for all workers to finish
		close(wp.results) // Close results channel
	})
}
