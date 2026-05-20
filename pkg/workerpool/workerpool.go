package workerpool

import (
	"context"
	"fmt"
	"sync"
)

// Job represents a unit of work
type Job struct {
	ID   int
	Data interface{}
}

// Result represents the output of processing a Job
type Result struct {
	Job    Job
	Output interface{}
	Err    error
}

// WorkerPool manages a pool of goroutines that process jobs
type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.RWMutex
	closed     bool
	stopOnce   sync.Once
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, jobQueueSize),
		results:    make(chan Result, jobQueueSize),
	}
}

// Start initializes the worker goroutines
func (wp *WorkerPool) Start(ctx context.Context, processor func(Job) Result) {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for {
				select {
				case job, ok := <-wp.jobs:
					if !ok {
						return // Channel closed, worker exits
					}
					result := processor(job)
					wp.results <- result
				case <-ctx.Done():
					return // Context cancelled, worker exits
				}
			}
		}(i)
	}
}

// Submit adds a job to the queue
func (wp *WorkerPool) Submit(job Job) error {
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
func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool) Stop() {
	wp.stopOnce.Do(func() {
		wp.mu.Lock()
		wp.closed = true
		close(wp.jobs) // No more jobs accepted
		wp.mu.Unlock()

		wp.wg.Wait()      // Wait for all workers to finish
		close(wp.results) // Close results channel
	})
}
