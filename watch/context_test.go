package watch

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"gotest.tools/assert"
)

func TestWatcherContext(t *testing.T) {
	var err error
	var ctx context.Context
	var watcherName string
	target := "WatcherContext"

	t.Logf("Given the need to test %s.", target)
	{
		t.Logf("\tTest 0:\tWhen calling %v.", target)
		{
			watcherName = gofakeit.UUID()
			ctx, err = WatcherContext(context.TODO(), watcherName, ConfigOption{ID: Units, Val: Microseconds}, ConfigOption{ID: LogFunc, Val: LogMetrics})

			assert.Assert(t, err == nil)
			t.Logf("\t%s\tShould be able to create WatcherContext with no error.", succeed)

			assert.Assert(t, ctx != nil)
			t.Logf("\t%s\tShould be able to create WatcherContext not nil returned.", succeed)

			assert.Equal(t, true, GetWatch(ctx).Running())
			t.Logf("\t%s\tThe context watch should have been running after creation %v.", succeed, target)
		}
	}
}

func TestWatcherContextOptionsValid(t *testing.T) {
	var err error
	var ctx context.Context
	var watcherName string
	target := "WatcherContext"

	t.Logf("Given the need to test %s with ConfigOptions.", target)
	{
		for i, unit := range []DurationUnits{Nanoseconds, Microseconds, Milliseconds, Seconds, Hours} {
			t.Logf("\tTest %d:\tWhen calling %v with valid ConfigOptions with valid Units <%v>.", i, target, unit)
			{
				watcherName = gofakeit.UUID()
				ctx, err = WatcherContext(context.TODO(), watcherName, ConfigOption{ID: Units, Val: unit})

				assert.Assert(t, err == nil)
				t.Logf("\t%s\tShould be able to create WatcherContext with no error.", succeed)

				assert.Equal(t, unit, GetWatch(ctx).Units())
				t.Logf("\t%s\tThe context watch should have been running after creation %v.", succeed, target)
			}
		}
	}
}
