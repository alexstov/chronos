// Package throt provides goroutine concurrency API for pooling and rate limiting.
package throt

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
)

// WorkerPool defines abstract worker pool of goroutines.
type WorkerPool interface {
	io.Closer
	Run(routine Routine)
	Size() int
	Completed() int
	Terminate() error
}

// NewWorkerPool creates a worker pool with a specific worker size.
func NewWorkerPool(ctx context.Context, workers int, limiter Limiter) WorkerPool {
	return newWorkerPool(ctx, workers, limiter)
}

// newWorkerPool creates a worker pool with a specific worker size.
func newWorkerPool(ctx context.Context, workers int, limiter Limiter) *workerPool {
	// There must be at least one worker.
	if workers < 1 {
		workers = 1
	}

	pool := workerPool{
		workers:  workers,
		routines: make(chan Routine),
		wg:       &sync.WaitGroup{},
	}
	pool.addWorkers(ctx, limiter)

	return &pool
}

func (p *workerPool) Size() int {
	return p.workers
}

func (p *workerPool) Completed() int {
	return int(p.completed)
}

// workerPool model.
type workerPool struct {
	workers   int
	completed int32
	routines  chan Routine
	wg        *sync.WaitGroup
	closed    atomic.Bool
}

// Run schedule the next job for the routines.
func (p *workerPool) Run(routine Routine) {
	p.routines <- routine
}

// Close waits for all goroutines to terminate (implements io.Closer).
func (p *workerPool) Close() error {
	return p.close(false)
}

// Terminate terminates all workers immediately.
func (p *workerPool) Terminate() error {
	return p.close(true)
}

// close waits for all goroutines to terminate (implements io.Closer).
func (p *workerPool) close(terminate bool) error {
	if p == nil {
		return nil
	}

	if p.closed.CompareAndSwap(false, true) {
		close(p.routines)
	}

	if !terminate {
		p.wg.Wait()
	}

	return nil
}

func (p *workerPool) addWorkers(ctx context.Context, limiter Limiter) {
	p.wg.Add(p.workers)

	for i := 0; i < p.workers; i++ {
		go func() {
			// Defer worker finishing all requests or cancelled.
			defer p.wg.Done()

			for {
				select {
				case routine, ok := <-p.routines:
					if !ok {
						return
					}

					// Throttle if a valid limiter exists.
					if limiter != nil {
						limiter.Await()
					}

					routine.Run()
					atomic.AddInt32(&p.completed, 1)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}
