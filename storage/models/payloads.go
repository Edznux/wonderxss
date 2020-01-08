package models

import (
	"time"
)

type Payload struct {
	ID          string
	Name        string
	Content     string
	ContentType string
	Hashes      SRIHashes
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
