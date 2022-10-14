package watch

// Optioner interface.
type Optioner interface {
	Option() Option
	Value() interface{}
}
