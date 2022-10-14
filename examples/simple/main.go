package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexstov/chronos/pkg/watch"
)

func main() {
	ctx, err := watch.WatcherContext(context.Background(), ChronosSampleSimple.String(), watch.ConfigOption{ID: watch.Units, Val: watch.Microseconds}, watch.ConfigOption{ID: watch.LogFunc, Val: LogMetrics})
	if err != nil {
		fmt.Printf("error")
	} else {
		defer func(ctx context.Context) { err = watch.Finish(ctx) }(ctx)
	}

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	t1, err := watch.Start(ctx, "Loop")
	if err != nil {
		fmt.Println("watch.Start(ctx, \"Loop\") failed")
	}
	for _, v := range xs {
		time.Sleep(time.Duration(v) * time.Microsecond)
		total += v
	}
	duration, _ := t1.Stop()
	fmt.Println("Loop duration is ", duration)

	t2, _ := watch.Start(ctx, OpsIO.String())
	fmt.Println(total / float64(len(xs)))
	t2.Stop()
}
