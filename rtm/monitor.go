package rtm

import (
	"time"

	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Monitor measures elapsed time in the area
type Monitor struct {
	ID      guuid.UUID
	Area    string
	elapsed time.Duration
	start   time.Time
	running bool
}

// NewMonitor creates new operation monitor
func NewMonitor(area string) *Monitor {
	var err error
	var monID guuid.UUID

	if monID, err = guuid.NewRandom(); err != nil {
		log.Errorf("RTM NewRandom failed to generate Monitor ID for <%v>.", area)
		return nil
	}

	return &Monitor{ID: monID, Area: area, running: false, elapsed: 0}
}

// Start starts the monitor
func (monitor *Monitor) Start() {
	if monitor == nil {
		return
	}

	if monitor.running {
		log.Warnf("RTM <%v> with ID <%v> is already running, get elapsed and restart.", monitor.Area, monitor.ID)
		monitor.Elapsed()
	}
	monitor.running = true
	monitor.start = time.Now()
}

// Elapsed adds to elapsed time.
func (monitor *Monitor) Elapsed() time.Duration {
	if !monitor.running {
		log.Errorf("RTM <%v> with ID <%v> hasn't been started.", monitor.Area, monitor.ID)
		return 0
	}
	elapsed := time.Since(monitor.start)
	monitor.running = false

	// Add to cumulative total
	monitor.elapsed += elapsed

	return elapsed
}

// Nanoseconds returns elapsed time in nanoseconds (ns)
func (monitor *Monitor) Nanoseconds() int64 {
	return monitor.elapsed.Nanoseconds()
}

// Microseconds returns elapsed time in 0.000001 of a second (μs/us)
func (monitor *Monitor) Microseconds() int64 {
	return monitor.elapsed.Microseconds()
}

// Milliseconds returns elapsed time in 0.001 of a second (ms)
func (monitor *Monitor) Milliseconds() int64 {
	return monitor.elapsed.Milliseconds()
}

// Seconds returns elapsed time as floating point number like 1.5s
func (monitor *Monitor) Seconds() float64 {
	return monitor.elapsed.Seconds()
}

// Hours returns elapsed time as floating point number like 0.23h
func (monitor *Monitor) Hours() float64 {
	return monitor.elapsed.Hours()
}
