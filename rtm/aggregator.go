package rtm

import (
	"time"
)

const aggrInitCapacity = 20

// Aggregator stores multiple monitors for the same operation
type Aggregator struct {
	Area     string
	Elapsed  time.Duration
	Monitors []*Monitor
}

// NewAggregator creates new area aggregator
func NewAggregator(area string) *Aggregator {
	return &Aggregator{Area: area, Monitors: make([]*Monitor, 0, aggrInitCapacity)}
}

// Add appends a monitor to the aggregator
func (aggr *Aggregator) Add(mon *Monitor) {
	aggr.Monitors = append(aggr.Monitors, mon)
}

// Aggregate aggregates elapsed durations of each monitor
func (aggr *Aggregator) Aggregate() {
	aggr.Elapsed = 0

	for _, mon := range aggr.Monitors {
		aggr.Elapsed += mon.elapsed
	}
}
