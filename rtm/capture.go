package rtm

import (
	"bytes"
	"fmt"
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Units of duration
type Units int

const (
	// Nanoseconds ...
	Nanoseconds Units = iota
	// Microseconds ...
	Microseconds
	// Milliseconds ...
	Milliseconds
	// Seconds ...
	Seconds
	// Minutes ...
	Minutes
	// Hours ...
	Hours
)

// TODO: move these to config
const totalArea = "Total"
const unknownArea = "Unknown"
const initMonitorSize = 0

// Capture benchmarks API Capture
type Capture struct {
	Name        string
	ID          guuid.UUID
	elapsed     *Monitor
	Unknown     time.Duration
	aggregators map[string]*Aggregator
	units       Units
}

// BeginCapture starts a new capture.
func BeginCapture(name string, args ...interface{}) (cap *Capture, err error) {
	var capID guuid.UUID

	if capID, err = guuid.NewRandom(); err != nil {
		log.Errorf("RTM NewRandom failed to generate Capture ID for <%v>.", name)
		return nil, err
	}

	c := Capture{Name: name, ID: capID, elapsed: NewMonitor(totalArea), aggregators: make(map[string]*Aggregator), units: Milliseconds}
	c.elapsed.Start()

	return &c, err
}

// Finish captures elapsed monitors, logs Capture info, and ends it.
func (cap *Capture) Finish() {
	for _, a := range cap.aggregators {
		a.Aggregate()
	}
	var aggrTotal time.Duration
	for _, a := range cap.aggregators {
		aggrTotal += a.Elapsed
	}
	cap.elapsed.Elapsed()

	cap.Unknown = cap.elapsed.elapsed - aggrTotal

	cap.Log()
}

// Start starts RTM area capture
func (cap *Capture) Start(sector string) (mon *Monitor) {
	if cap == nil {
		return
	}

	var aggr *Aggregator
	var has bool

	// TODO: start immediately, add in another goroutine, signal
	mon = NewMonitor(sector) // TODO: NewMonitor returns pre-created monitor?
	if aggr, has = cap.aggregators[sector]; !has {
		aggr = NewAggregator(sector) // TODO: NewAggregator returns pre-created aggregator?
		cap.aggregators[sector] = aggr
	}
	aggr.Add(mon)
	mon.Start()

	return mon
}

// Log logs Capture monitors as Info message.
func (cap *Capture) Log() {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "RTM %v, ID <%v>: Total <%v>, ", cap.Name, cap.ID, durationString(cap.elapsed.elapsed, cap.units))

	for _, a := range cap.aggregators {
		buf.WriteString(a.Area)
		buf.WriteString(" <")
		buf.WriteString(durationString(a.Elapsed, cap.units))
		buf.WriteString(">, ")
	}
	buf.Truncate(buf.Len() - 2)
	buf.WriteString(".")

	log.Info(buf.String())
}

func durationString(duration time.Duration, units Units) string {
	switch units {
	case Nanoseconds:
		return strconv.FormatInt(duration.Nanoseconds(), 10) + "ns"
	case Microseconds:
		return strconv.FormatInt(duration.Microseconds(), 10) + "us"
	case Milliseconds:
		return strconv.FormatInt(duration.Milliseconds(), 10) + "ms"
	case Seconds:
		return strconv.FormatFloat(duration.Seconds(), 'f', 6, 64) + "s"
	case Minutes:
		return strconv.FormatFloat(duration.Minutes(), 'f', 6, 64) + "min"
	case Hours:
		return strconv.FormatFloat(duration.Hours(), 'f', 6, 64) + "h"
	default:
		return strconv.FormatInt(duration.Milliseconds(), 10) + "ms"
	}
}
