package watch

// Option of the capture.
type Option int

const (
	// Units to output monitor time i.e. nanoseconds(ns). etc.
	Units Option = iota
	// LogFunc function to log watch output.
	LogFunc
)

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

// ConfigOption of th capture.
type ConfigOption struct {
	ID  Option
	Val interface{}
}

func (o ConfigOption) Option() Option {
	return o.ID
}

func (o ConfigOption) Value() interface{} {
	return o.Val
}
