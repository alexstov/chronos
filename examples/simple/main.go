package main

import (
	"context"
	"fmt"
	"time"

	watch2 "github.com/alexstov/chronos/watch"
)

func main() {
	ctx, err := watch2.WatcherContext(context.Background(), ChronosSampleSimple.String(), watch2.ConfigOption{ID: watch2.Units, Val: watch2.Microseconds}, watch2.ConfigOption{ID: watch2.LogFunc, Val: LogMetrics})
	if err != nil {
		fmt.Printf("error")
	} else {
		defer func(ctx context.Context) { err = watch2.Finish(ctx) }(ctx)
	}

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	m1, err := watch2.Start(ctx, "Loop")
	if err != nil {
		fmt.Println("watch.Start(ctx, \"Loop\") failed")
	}
	for _, v := range xs {
		time.Sleep(time.Duration(v) * time.Millisecond)
		total += v
	}
	duration, _ := m1.Stop()

	m2, _ := watch2.Start(ctx, OpsIO.String())
	fmt.Println("Loop duration is ", duration)
	fmt.Println("Total sleep time in the loop is ", time.Duration(total)*time.Millisecond)
	_, _ = m2.Stop()

	m3, _ := watch2.NewMonitor("Untracked", true)
	time.Sleep(time.Duration(1500) * time.Millisecond)
	t3, _ := m3.Stop()
	fmt.Printf("Untracked delay <%v>\n", t3)
}
