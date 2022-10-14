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
