package locker

// Locker with distribute
type Locker interface {
	Lock() error
	Finish() error
	Heartbeat() error
	IsLocked() (bool, error)
}
