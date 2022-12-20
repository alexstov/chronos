package watch

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestWatchNilPtr(t *testing.T) {
	t.Log("Given the need to test Start when nil Monitor pointer is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v using nil Capture poiter.", "m.Start")
		{
			var watch *Watch
			ctx := context.Background()
			name := gofakeit.UUID()
			watch.Start(ctx, name)
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}

func TestWatchNilContext(t *testing.T) {
	t.Log("Given the need to test Start when nil Monitor pointer is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v using nil Capture poiter.", "m.Start")
		{
			var watch *Watch
			var ctx context.Context
			name := gofakeit.UUID()

			watch = &Watch{}
			watch.Start(ctx, name)
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}
