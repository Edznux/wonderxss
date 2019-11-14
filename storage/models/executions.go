package models

import "time"

// Execution represents an execution of a payload.
// It is unique. A new trigger will generate a new Execution.
type Execution struct {
	ID          string
	PayloadID   string
	AliasID     string
	TriggeredAt time.Time
}
