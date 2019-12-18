package models

import "time"

type Injection struct {
	ID         string
	Name       string
	Content    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
