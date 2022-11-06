package throt

import "time"

type Limiter interface {
	// Await blocks execution for the desired throttle.
	Await()
}

// NewLimiter creates a new limiter.
func NewLimiter(d time.Duration, count int) Limiter {
	l := &limiter{
		maxCount: count,
		count:    count,
		ticker:   time.NewTicker(d),
		ch:       make(chan struct{}),
	}
	go l.run()

	return l
}

// limiter throttles number of requests per duration interval.
type limiter struct {
	maxCount int
	count    int
	ticker   *time.Ticker
	ch       chan struct{}
}

func (l *limiter) run() {
	if l == nil {
		return
	}

	for {
		// if counter has reached 0: block until next tick.
		if l.count <= 0 {
			<-l.ticker.C
			l.count = l.maxCount
		}

		// otherwise:
		select {
		case l.ch <- struct{}{}:
			l.count--

		case <-l.ticker.C:
			l.count = l.maxCount
		}
	}
}

// Await blocks execution for the desired throttle.
func (l *limiter) Await() {
	if l == nil {
		return
	}

	<-l.ch
}
