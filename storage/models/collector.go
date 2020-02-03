package models

import "time"

// Loot represents an Loot of a payload.
// It is unique. A new trigger will generate a new Loot.
type Loot struct {
	ID        string
	Data      string
	CreatedAt time.Time
}
