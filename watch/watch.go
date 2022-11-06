// Package watch provides performance monitoring routines
package watch

import (
	"context"
	"strconv"
	"time"

	"github.com/gildas/go-errors"
)

const (
	unknownArea = "Unknown"
	totalArea   = "Total"
)

// NewWatcher creates a new watcher.
func NewWatcher(name string, options ...interface{}) (watch *Watch, err error) {
	m, err := NewMonitor(totalArea, false)
	if err != nil {
		return nil, err
	}

	w := Watch{
		Name:        name,
		elapsed:     m,
		aggregators: make(map[string]*Aggregator),
		units:       Milliseconds,
		percentage:  DecimalPlacesOnes,
		precision:   DecimalPlacesMillions,
	}

	w.setWatcherOptions(options)

	if err = w.elapsed.Start(); err != nil {
		return nil, err
	}

	return &w, err
}

// LogMetricsFunc defines a function to log the watch metrics.
type LogMetricsFunc func(kvp [][]string)

// Watch performance profiler instance.
type Watch struct {
	context     context.Context
	Name        string
	elapsed     *Monitor
	Unknown     time.Duration
	aggregators map[string]*Aggregator
	units       DurationUnits
	LogMetrics  LogMetricsFunc
	percentage  DecimalPlaces
	precision   DecimalPlaces
}

// Start starts a new monitor.
func Start(ctx context.Context, sector string) (mon *Monitor, err error) {
	watch := GetWatch(ctx)
	if watch == nil {
		return nil, errors.ArgumentInvalid.With("ctx", ctx)
	}
	return watch.Start(ctx, sector), nil
}

// Finish captures elapsed monitors, logs Capture info, and ends it.
func Finish(ctx context.Context) (err error) {
	w := GetWatch(ctx)
	return w.Finish(ctx)
}

// Running returns true if watch is running.
func (w *Watch) Running() bool {
	if w == nil {
		return false
	}
	return w.elapsed.running
}

// Start starts RTM area capture.
func (w *Watch) Start(ctx context.Context, sector string) (mon *Monitor) {
	if w == nil {
		return nil
	}
	if w.aggregators == nil {
		return nil
	}
	w.context = ctx

	var err error
	var aggr *Aggregator
	var has bool

	// TODO: start immediately, add in another goroutine, signal
	// TODO: NewMonitor returns pre-created monitor?
	if mon, err = NewMonitor(sector, false); err != nil {
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
func (w *Watch) Finish(ctx context.Context) (err error) {
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
		{"Trace", GetTrace(w.context).String()},
		{totalArea, durationString(w.elapsed.elapsed, w.units, w.precision), percentageString(w.elapsed.elapsed, w.elapsed.elapsed, w.percentage)},
	}

	for _, a := range w.aggregators {
		fields = append(fields, []string{a.Area, durationString(a.Elapsed, w.units, w.precision), percentageString(a.Elapsed, w.elapsed.elapsed, w.percentage)})
	}

	fields = append(fields, []string{unknownArea, durationString(w.Unknown, w.units, w.precision), percentageString(w.Unknown, w.elapsed.elapsed, w.percentage)})

	if w.LogMetrics != nil {
		w.LogMetrics(fields)
	}
}

func (w *Watch) setWatcherOptions(options ...interface{}) {
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
						case Percentage:
							w.percentage = ops.Value().(DecimalPlaces)
						case Precision:
							w.precision = ops.Value().(DecimalPlaces)
						}
					}
				}
			}
		}
	}
}

func (w *Watch) Units() DurationUnits {
	return w.units
}

func percentageString(elapsed time.Duration, total time.Duration, precision DecimalPlaces) string {
	prec := int(precision) - 1
	percentage := float64(elapsed) / float64(total) * 100

	return strconv.FormatFloat(percentage, 'f', prec, 64) + "%"
}

func durationString(duration time.Duration, units DurationUnits, precision DecimalPlaces) string {
	const baseTen = 10
	const numberSixtyFour = 64
	prec := int(precision) - 1

	switch units {
	case Nanoseconds:
		return strconv.FormatInt(duration.Nanoseconds(), baseTen) + "ns"
	case Microseconds:
		return strconv.FormatInt(duration.Microseconds(), baseTen) + "µs"
	case Milliseconds:
		return strconv.FormatInt(duration.Milliseconds(), baseTen) + "ms"
	case Seconds:
		return strconv.FormatFloat(duration.Seconds(), 'f', prec, numberSixtyFour) + "s"
	case Minutes:
		return strconv.FormatFloat(duration.Minutes(), 'f', prec, numberSixtyFour) + "min"
	case Hours:
		return strconv.FormatFloat(duration.Hours(), 'f', prec, numberSixtyFour) + "h"
	default:
		return strconv.FormatInt(duration.Milliseconds(), baseTen) + "ms"
	}
}
