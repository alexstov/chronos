package watch

type Optioner interface {
	Option() Option
	Value() interface{}
}
