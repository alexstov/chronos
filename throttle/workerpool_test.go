package throt

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTask struct {
	name string
	fn   func()
}

func (t testTask) Run() {
	t.fn()
}

func TestWorkerPool(t *testing.T) {
	ctx := context.Background()
	target := "NewWorkerPool"
	tasksNum := 200

	t.Logf("Given the need to test %s.", target)
	pool := newPoolRun(ctx, 100, tasksNum, nil)
	t.Logf("\t%s\tAfter calling pool.Close()", succeed)
	_ = pool.Close()
	t.Logf("\t%s\tShould have all tasks compleated.", succeed)
	assert.Equal(t, tasksNum, pool.Completed())
}

func TestWorkerPool_Close(t *testing.T) {
	ctx := context.Background()

	target := "NewWorkerPool Close"
	t.Logf("Given the need to test %s.", target)

	pool := newPoolRun(ctx, 100, 200, nil)
	err1 := pool.Close()
	err2 := pool.Close()
	t.Logf("\t%s\tAfter calling pool.Close() twice", succeed)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	t.Logf("\t%s\tBoth calls should return nil error", succeed)
}

func newPoolRun(ctx context.Context, workersNum, tasksNum int, limiter Limiter) (pool *workerPool) {
	var wg sync.WaitGroup
	var scheduledNum int32

	pool = newWorkerPool(ctx, workersNum, limiter)

	wg.Add(tasksNum)
	var tasks []testTask
	for i := 0; i < tasksNum; i++ {
		wg.Done()
		tasks = append(tasks, testTask{
			fmt.Sprintf("task #%d", i),
			func() {
				atomic.AddInt32(&scheduledNum, 1)
			},
		})
	}

	for _, task := range tasks {
		pool.Run(task)
	}

	return pool
}
