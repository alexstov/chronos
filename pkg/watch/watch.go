// Package chronos provides performance monitoring routines
package watch

import (
	"context"
	"strconv"
	"time"
)

// NewWatcher creates a new watcher.
func NewWatcher(name string, options ...interface{}) (watch *Watch, err error) {
	m, err := NewMonitor(totalArea)
	if err != nil {
		return nil, err
	}

	w := Watch{
		Name:        name,
		elapsed:     m,
		aggregators: make(map[string]*Aggregator),
		units:       Milliseconds,
	}

	w.setCaptureOptions(options)

	if err = w.elapsed.Start(); err != nil {
		return nil, err
	}

	return &w, err
}

// DurationUnits of duration.
type DurationUnits int

const (
	// Nanoseconds ...
	Nanoseconds DurationUnits = iota
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

const (
	unknownArea     = "Unknown"
	totalArea       = "Total"
	initMonitorSize = 0
)

// LogMetricsFunc defines a function to log the watch metrics.
type LogMetricsFunc func(kvp [][]string)

// Watch performance profiler instance.
type Watch struct {
	Name        string
	elapsed     *Monitor
	Unknown     time.Duration
	aggregators map[string]*Aggregator
	units       DurationUnits
	LogMetrics  LogMetricsFunc
	context     context.Context
}

// WatcherContext adds watcher to the context.
func WatcherContext(ctx context.Context, name string, options ...Optioner) (context.Context, error) {
	watch, err := NewWatcher(name, options)
	if err != nil {
		return nil, err
	}

	watchCtx := NewTraceWatchContext(ctx, watch)

	return watchCtx, nil
}

// Start starts a new monitor.
func Start(ctx context.Context, sector string) (mon *Monitor, err error) {
	watch := GetWatch(ctx)
	watch.context = ctx
	return watch.start(sector), nil
}

// Start starts RTM area capture.
func (w *Watch) start(sector string) (mon *Monitor) {
	if w == nil {
		return
	}

	var err error
	var aggr *Aggregator
	var has bool

	// TODO: start immediately, add in another goroutine, signal
	// TODO: NewMonitor returns pre-created monitor?
	if mon, err = NewMonitor(sector); err != nil {
		return nil
	}

	if aggr, has = w.aggregators[sector]; !has {
		aggr = NewAggregator(sector) // TODO: NewAggregator returns pre-created aggregator?
		w.aggregators[sector] = aggr
	}
	aggr.Add(mon)
	_ = mon.Start()

	return mon
}

// Finish captures elapsed monitors, logs Capture info, and ends it.
func Finish(ctx context.Context) (err error) {
	w := GetWatch(ctx)

	for _, a := range w.aggregators {
		a.Aggregate()
	}
	var aggrTotal time.Duration
	for _, a := range w.aggregators {
		aggrTotal += a.Elapsed
	}

	if _, err = w.elapsed.Stop(); err != nil {
		return err
	}

	w.Unknown = w.elapsed.elapsed - aggrTotal

	w.log()

	return nil
}

func (w *Watch) log() {
	fields := [][]string{
		{"Watcher", w.Name},
		{"Total", durationString(w.elapsed.elapsed, w.units)},
	}

	for _, a := range w.aggregators {
		fields = append(fields, []string{a.Area, durationString(a.Elapsed, w.units)})
	}

	fields = append(fields, []string{"Unknown", durationString(w.Unknown, w.units)})

	if w.LogMetrics != nil {
		w.LogMetrics(fields)
	}
}

func (w *Watch) setCaptureOptions(options ...interface{}) {
	for _, slice := range options {
		iSlice, ok := slice.([]interface{})
		if ok {
			for _, o := range iSlice {
				opsSlice, ok := o.([]Optioner)
				if ok {
					for _, ops := range opsSlice {
						switch ops.Option() {
						case Units:
							w.units, _ = ops.Value().(DurationUnits)
						case LogFunc:
							w.LogMetrics = ops.Value().(func([][]string))
						}
					}
				}
			}
		}
	}
}

func durationString(duration time.Duration, units DurationUnits) string {
	const baseTen = 10
	const numberSix = 6
	const numberSixtyFour = 64

	switch units {
	case Nanoseconds:
		return strconv.FormatInt(duration.Nanoseconds(), baseTen) + "ns"
	case Microseconds:
		return strconv.FormatInt(duration.Microseconds(), baseTen) + "µs"
	case Milliseconds:
		return strconv.FormatInt(duration.Milliseconds(), baseTen) + "ms"
	case Seconds:
		return strconv.FormatFloat(duration.Seconds(), 'f', numberSix, numberSixtyFour) + "s"
	case Minutes:
		return strconv.FormatFloat(duration.Minutes(), 'f', numberSix, numberSixtyFour) + "min"
	case Hours:
		return strconv.FormatFloat(duration.Hours(), 'f', numberSix, numberSixtyFour) + "h"
	default:
		return strconv.FormatInt(duration.Milliseconds(), baseTen) + "ms"
	}
}
