package models

import "time"

// Alias represent a translation from a Payload ID to a shorter alias
// Useful for short injection and easier to remember
// E.g: /p/5c28e6cf-72d6-49b3-b03e-2a40672c41ad would be shortened /p/a
// It also works for subdomains like a.domain.tld
type Alias struct {
	ID         string
	PayloadID  string
	Short      string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
