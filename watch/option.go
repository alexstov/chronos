package watch

// ConfigOption of th capture.
type ConfigOption struct {
	ID  Option
	Val interface{}
}

// Option return config option identifier.
func (o ConfigOption) Option() Option {
	return o.ID
}

// Value return config option value.
func (o ConfigOption) Value() interface{} {
	return o.Val
}

// Option of the capture.
type Option int

const (
	// Units to output monitor time i.e. nanoseconds(ns). etc.
	Units Option = iota
	// LogFunc function to log watch output.
	LogFunc
	// Percentage option for performance stats.
	Percentage
	// Precision option for float values.
	Precision
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

// DecimalPlaces for percentage and unit values.
type DecimalPlaces int

const (
	// DecimalPlacesAll disabled.
	DecimalPlacesAll DecimalPlaces = iota
	// DecimalPlacesOnes 1 - ones place.
	DecimalPlacesOnes
	// DecimalPlacesTenths 2 - tenths place.
	DecimalPlacesTenths
	// DecimalPlacesHundredths 3 - hundredths place.
	DecimalPlacesHundredths
	// DecimalPlacesThousandths 4 - thousandths place.
	DecimalPlacesThousandths
	// DecimalPlacesTenThousandths 5 - ten thousandths place.
	DecimalPlacesTenThousandths
	// DecimalPlacesHundredThousandths  6 - hundred thousandths place.
	DecimalPlacesHundredThousandths
	// DecimalPlacesMillions 7 - millions place.
	DecimalPlacesMillions
)
