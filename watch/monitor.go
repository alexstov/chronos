package watch

import (
	"time"

	"github.com/gildas/go-errors"
	"github.com/google/uuid"
)

// Monitor measures elapsed time in the area.
type Monitor struct {
	ID      uuid.UUID
	Area    string
	elapsed time.Duration
	start   time.Time
	running bool
}

// NewMonitor creates new operation monitor.
func NewMonitor(area string, start bool) (*Monitor, error) {
	var err error
	var monID uuid.UUID

	if monID, err = uuid.NewRandom(); err != nil {
		return nil, errors.Wrapf(err, "failed to create the new monitor Random ID for <%v>", area)
	}

	monitor := Monitor{ID: monID, Area: area, running: false, elapsed: 0}

	if start {
		monitor.Start()
	}

	return &monitor, nil
}

// Start starts a new monitor.
func (m *Monitor) Start() (err error) {
	if m == nil {
		err := errors.Error{}
		err.WithCause(errors.NotInitialized.With("monitor is nil"))
		return err
	}

	if m.running {
		if _, err = m.Stop(); err != nil {
			return err
		}
	}
	m.running = true
	m.start = time.Now()

	return nil
}

// Stop adds to monitor elapsed time and stops time capture.
func (m *Monitor) Stop() (time.Duration, error) {
	if m == nil {
		return 0, nil
	}

	if !m.running {
		return 0, errors.RuntimeError.With("cannot stop already stopped monitor", m)
	}
	elapsed := time.Since(m.start)
	m.running = false

	// Add to cumulative total
	m.elapsed += elapsed

	return elapsed, nil
}

// Nanoseconds returns elapsed time in nanoseconds (ns).
func (m *Monitor) Nanoseconds() int64 {
	return m.elapsed.Nanoseconds()
}

// Microseconds returns elapsed time in 0.000001 of a second (μs/us).
func (m *Monitor) Microseconds() int64 {
	return m.elapsed.Microseconds()
}

// Milliseconds returns elapsed time in 0.001 of a second (ms).
func (m *Monitor) Milliseconds() int64 {
	return m.elapsed.Milliseconds()
}

// Seconds returns elapsed time as floating point number like 1.5s.
func (m *Monitor) Seconds() float64 {
	return m.elapsed.Seconds()
}

// Hours returns elapsed time as floating point number like 0.23h.
func (m *Monitor) Hours() float64 {
	return m.elapsed.Hours()
}
