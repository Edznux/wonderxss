package models

type StorageError uint

const (
	Success StorageError = iota
	NoSuchItem
	AlreadyExist
	StorageFailure
)

func (s StorageError) Error() string {
	switch s {
	case Success:
		return "OK"
	case NoSuchItem:
		return "No such item"
	case AlreadyExist:
		return "Already exists"
	case StorageFailure:
		return "The data store has encountered an unexpected failure"
	default:
		return "Unknown Code (The developer forgot to add it to the String() switch ?)"
	}
}
