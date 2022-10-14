package watch

// Watcher interface.
type Watcher interface {
	Stop() (err error)
}
