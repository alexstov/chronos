package main

import (
	"context"
	"fmt"
	"github.com/alexstov/chronos/pkg/watch"
)

func main() {
	ctx, err := watch.WatcherContext(context.Background(), ChronosSampleSimple.String(), watch.ConfigOption{ID: watch.Units, Val: watch.Microseconds}, watch.ConfigOption{ID: watch.LogFunc, Val: LogMetrics})
	if err != nil {
		fmt.Printf("error")
	} else {
		defer watch.Finish(ctx)
	}

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	t1, _ := watch.Start(ctx, "Loop")
	for _, v := range xs {
		//time.Sleep(time.Duration(v) * time.Nanosecond)
		total += v
	}
	t1.Stop()

	t2, _ := watch.Start(ctx, OpsIO.String())
	fmt.Println(total / float64(len(xs)))
	t2.Stop()
}
