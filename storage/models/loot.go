package models

import "time"

// Loot represents an execution of a payload.
// It is unique. A new trigger will generate a new Loot.
type Loot struct {
	ID          string
	PayloadID   string
	AliasID     string
	TriggeredAt time.Time
}
