package rtm

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// OpsID operation ID
type OpsID int

const (
	// OpsTotal unknown operation
	OpsTotal OpsID = iota
	// OpsUnknown unknown operation
	OpsUnknown
	// OpsCritical ...
	OpsCritical
	// OpsCustomer ...
	OpsCustomer
	// OpsSendEmail ...
	OpsSendEmail
	// OpsMongoDB ...
	OpsMongoDB
	// OpsAddress ...
	OpsAddress
	// OpsWeDeliver ...
	OpsWeDeliver
	// OpsPayment ...
	OpsPayment
)

func (o OpsID) String() string {
	return [...]string{"Total", "Unknown", "Customer", "SendEmail", "MonogDB", "Address", "WeDeliver", "Payment"}[o]
}

// Monitor measures operation elapsed time
type Monitor struct {
	ID      OpsID
	elapsed time.Duration
	start   time.Time
	running bool
}

// NewMonitor creates new operation monitor
func NewMonitor(opsID OpsID) *Monitor {
	return &Monitor{ID: opsID, running: false, elapsed: 0}
}

// Start starts the monitor
func (monitor *Monitor) Start(opsID OpsID) {
	if monitor == nil {
		return
	}

	if monitor.running {
		log.Warnf("RTM monitor %v is already running, get elapsed and restart.", opsID.String())
		monitor.Elapsed()
	}
	monitor.running = true
	monitor.start = time.Now()
}

// Elapsed adds to elapsed time.
func (monitor *Monitor) Elapsed() {
	if !monitor.running {
		log.Errorf("RTM monitor %v hasn't been started.", monitor.ID.String())
		return
	}
	elapsed := time.Since(monitor.start)
	monitor.running = false
	monitor.elapsed += elapsed
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
