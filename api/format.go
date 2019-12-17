package api

import (
	"time"

	"github.com/edznux/wonderxss/storage/models"
)

// APIError defines all the errors that can be sent by the API.
type APIError int

const (
	Success APIError = iota
	NotFound
	AlreadyExist
	DatabaseError
	InvalidInput
	MissingAuthorization
	InvalidAuthorization
	NotImplementedYet
)

func (s APIError) Error() string {
	switch s {
	case Success:
		return "OK"
	case NotFound:
		return "Not found"
	case AlreadyExist:
		return "Already exists"
	case DatabaseError:
		return "The database encoutered an unexpected error"
	case InvalidInput:
		return "Invalid input"
	case MissingAuthorization:
		return "Missing authorization"
	case InvalidAuthorization:
		return "Invalid authorization"
	case NotImplementedYet:
		return "This feature is not implemented yet"
	default:
		return "Unknown APIError Code  (The developer forgot to add it to the String() switch ?)"
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Payload represent the structure of the frontend-facing payload. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Payload struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	Content    string           `json:"content"`
	Hashes     models.SRIHashes `json:"hashes"` // Used for SRI (sub ressource integrity)
	CreatedAt  time.Time        `json:"created_at"`
	ModifiedAt time.Time        `json:"modified_at"`
}

func (p Payload) fromStorage(s models.Payload) Payload {
	p.ID = s.ID
	p.Name = s.Name
	p.Content = s.Content
	p.Hashes = s.Hashes
	p.CreatedAt = s.CreatedAt
	p.ModifiedAt = s.ModifiedAt
	return p
}

// Aliases represent the structure of the frontend-facing Aliases. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Alias struct {
	ID         string    `json:"id"`
	PayloadID  string    `json:"payload_id"` // Used for SRI (sub ressource integrity)
	Alias      string    `json:"alias"`      // Used for SRI (sub ressource integrity)
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (p Alias) fromStorage(s models.Alias) Alias {
	p.ID = s.ID
	p.PayloadID = s.PayloadID
	p.Alias = s.Short
	p.CreatedAt = s.CreatedAt
	p.ModifiedAt = s.ModifiedAt
	return p
}

// Execution represent the structure of the frontend-facing Executions. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Execution struct {
	ID          string    `json:"id"`
	PayloadID   string    `json:"payload_id"`
	AliasID     string    `json:"alias_id"`
	TriggeredAt time.Time `json:"triggered_at"`
}

func (l Execution) fromStorage(s models.Execution) Execution {
	l.ID = s.ID
	l.PayloadID = s.PayloadID
	l.AliasID = s.AliasID
	l.TriggeredAt = s.TriggeredAt
	return l
}

// Execution represent the structure of the frontend-facing Executions. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Collector struct {
	ID        string    `json:"id"`
	PayloadID string    `json:"payload_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func (l Collector) fromStorage(s models.Collector) Collector {
	l.ID = s.ID
	l.PayloadID = s.PayloadID
	l.CreatedAt = s.CreatedAt
	l.Data = s.Data
	return l
}
