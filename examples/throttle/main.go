package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexstov/chronos/throttle"
)

type simpleConsoleOutput struct {
	ctx  context.Context
	job  string
	task int
}

func (j simpleConsoleOutput) Run() {
	fmt.Printf("%s %d\n", j.job, j.task)
}

func bursts(l throt.Limiter) {
	ctx := context.Background()
	start := time.Now()
	workerPool := throt.NewWorkerPool(ctx, 5, l)
	defer func() { _ = workerPool.Close() }()

	for i := 0; i < 20; i++ {
		task := simpleConsoleOutput{
			ctx:  ctx,
			job:  "burst",
			task: i,
		}
		workerPool.Run(task)
	}

	fmt.Println("--- bursts:", time.Now().Sub(start), "elapsed")
}

func slowIterations(l throt.Limiter) {
	ctx := context.Background()
	start := time.Now()
	workerPool := throt.NewWorkerPool(ctx, 5, l)
	defer func() { _ = workerPool.Close() }()

	for i := 0; i < 20; i++ {
		<-time.After(250 * time.Millisecond)
		task := simpleConsoleOutput{
			ctx:  ctx,
			job:  "iteration",
			task: i,
		}
		workerPool.Run(task)
	}

	fmt.Println("--- slowIterations:", time.Now().Sub(start), "elapsed")
}

func main() {
	fmt.Println("Hello, 世界")

	bursts(throt.NewLimiter(time.Second, 2))
	slowIterations(throt.NewLimiter(time.Second, 3))
}
