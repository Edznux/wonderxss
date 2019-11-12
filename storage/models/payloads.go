package models

import (
	"time"
)

type Payload struct {
	ID         string
	Name       string
	Content    string
	Hash       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
