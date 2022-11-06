package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alexstov/chronos/watch"
)

func main() {
	ctx, err := watch.WatcherContext(context.Background(), ChronosSampleSimple.String(),
		watch.ConfigOption{ID: watch.Units, Val: watch.Seconds},
		watch.ConfigOption{ID: watch.Precision, Val: watch.DecimalPlacesHundredths},
		watch.ConfigOption{ID: watch.LogFunc, Val: LogMetrics},
		watch.ConfigOption{ID: watch.Percentage, Val: watch.DecimalPlacesTenths})
	defer func() { _ = watch.Finish(ctx) }()

	if err != nil {
		fmt.Printf("error")
	} else {
		defer func(ctx context.Context) { err = watch.Finish(ctx) }(ctx)
	}

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	m1, err := watch.Start(ctx, "Loop")
	if err != nil {
		fmt.Println("watch.Start(ctx, \"Loop\") failed")
	}
	for _, v := range xs {
		time.Sleep(time.Duration(v) * time.Millisecond)
		total += v
	}
	duration, _ := m1.Stop()

	m2, _ := watch.Start(ctx, OpsIO.String())
	fmt.Println("Loop duration is ", duration)
	fmt.Println("Total sleep time in the loop is ", time.Duration(total)*time.Millisecond)
	_, _ = m2.Stop()

	m3, _ := watch.NewMonitor("Untracked", true)
	time.Sleep(time.Duration(1500) * time.Millisecond)
	t3, _ := m3.Stop()
	fmt.Printf("Untracked delay <%v>\n", t3)
}
