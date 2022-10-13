package watch

// Watcher interface.
type Watcher interface {
	Finish() (err error)
}
