package watch

import "context"

// Watcher interface.
type Watcher interface {
	Start(ctx context.Context, sector string) *Monitor
	Finish(ctx context.Context) error
	Running() bool
	Units() DurationUnits
}
