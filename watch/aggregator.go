package watch

import (
	"github.com/gildas/go-errors"
	"time"
)

const aggrInitCapacity = 20

// Aggregator stores multiple monitors for the same operation.
type Aggregator struct {
	Area     string
	Elapsed  time.Duration
	Monitors []*Monitor
}

// NewAggregator creates new area aggregator.
func NewAggregator(area string) *Aggregator {
	return &Aggregator{Area: area, Monitors: make([]*Monitor, 0, aggrInitCapacity)}
}

// Add appends a monitor to the aggregator.
func (a *Aggregator) Add(mon *Monitor) {
	a.Monitors = append(a.Monitors, mon)
}

// Aggregate aggregates elapsed durations of each monitor.
func (a *Aggregator) Aggregate() (err error) {
	a.Elapsed = 0

	for _, mon := range a.Monitors {
		if mon.running {
			if _, err = mon.Stop(); err == nil {
				err = errors.RuntimeError.With("dangling monitors may distort aggregator accuracy", "monitor", mon.Area)
			}
		}
		a.Elapsed += mon.elapsed
	}

	return err
}
