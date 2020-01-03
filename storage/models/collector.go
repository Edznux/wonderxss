package models

import "time"

// Collector represents an Collector of a payload.
// It is unique. A new trigger will generate a new Collector.
type Collector struct {
	ID        string
	Data      string
	CreatedAt time.Time
}
