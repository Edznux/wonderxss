package api

import (
	"time"

	"github.com/edznux/wonderxss/storage/models"
)

// Error defines all the errors that can be sent by the API.
type Error int

// Enum all the possible error from the API
// They all have String() representation
const (
	Success Error = iota
	NotFound
	AlreadyExist
	DatabaseError
	InvalidInput
	MissingAuthorization
	InvalidAuthorization
	NotImplementedYet
)

func (s Error) Error() string {
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
		return "Unknown Error Code  (The developer forgot to add it to the String() switch ?)"
	}
}

// Response represent the main structure for ALL the response from the API
// The error must be empty if there wasn't an error
// They should also be only using the api.Error model
// Data can be anything but we try to not nest to many things
type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// Payload represent the structure of the frontend-facing payload. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type Payload struct {
	ID          string           `json:"id" mapstructure:"id"`
	Name        string           `json:"name" mapstructure:"name"`
	Content     string           `json:"content" mapstructure:"content"`
	ContentType string           `json:"content_type" mapstructure:"content_type"`
	Hashes      models.SRIHashes `json:"hashes" mapstructure:"hashes"` // Used for SRI (sub ressource integrity)
	CreatedAt   time.Time        `json:"created_at" mapstructure:"created_at"`
	ModifiedAt  time.Time        `json:"modified_at" mapstructure:"modified_at"`
}

func (p Payload) FromStorage(s models.Payload) Payload {
	p.ID = s.ID
	p.Name = s.Name
	p.Content = s.Content
	p.ContentType = s.ContentType
	p.Hashes = s.Hashes
	p.CreatedAt = s.CreatedAt
	p.ModifiedAt = s.ModifiedAt
	return p
}

// Aliases represent the structure of the frontend-facing Aliases. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type Alias struct {
	ID         string    `json:"id" mapstructure:"id"`
	PayloadID  string    `json:"payload_id" mapstructure:"payload_id"` // Used for SRI (sub ressource integrity)
	Alias      string    `json:"alias" mapstructure:"alias"`           // Used for SRI (sub ressource integrity)
	CreatedAt  time.Time `json:"created_at" mapstructure:"created_at"`
	ModifiedAt time.Time `json:"modified_at" mapstructure:"modified_at"`
}

func (p Alias) FromStorage(s models.Alias) Alias {
	p.ID = s.ID
	p.PayloadID = s.PayloadID
	p.Alias = s.Short
	p.CreatedAt = s.CreatedAt
	p.ModifiedAt = s.ModifiedAt
	return p
}

// Execution represent the structure of the frontend-facing Executions. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type Execution struct {
	ID          string    `json:"id" mapstructure:"id"`
	PayloadID   string    `json:"payload_id" mapstructure:"payload_id"`
	AliasID     string    `json:"alias_id" mapstructure:"alias_id"`
	TriggeredAt time.Time `json:"triggered_at" mapstructure:"triggered_at"`
}

func (l Execution) FromStorage(s models.Execution) Execution {
	l.ID = s.ID
	l.PayloadID = s.PayloadID
	l.AliasID = s.AliasID
	l.TriggeredAt = s.TriggeredAt
	return l
}

// Collector represent the structure of the frontend-facing Executions. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type Collector struct {
	ID        string    `json:"id" mapstructure:"id"`
	Data      string    `json:"data" mapstructure:"data"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
}

func (l Collector) FromStorage(s models.Collector) Collector {
	l.ID = s.ID
	l.CreatedAt = s.CreatedAt
	l.Data = s.Data
	return l
}

// Injection represent the structure of the frontend-facing Executions. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type Injection struct {
	ID         string    `json:"id" mapstructure:"id"`
	Name       string    `json:"name" mapstructure:"name"`
	Content    string    `json:"content" mapstructure:"content"`
	CreatedAt  time.Time `json:"created_at" mapstructure:"created_at"`
	ModifiedAt time.Time `json:"modified_at" mapstructure:"modified_at"`
}

func (l Injection) FromStorage(s models.Injection) Injection {
	l.ID = s.ID
	l.Name = s.Name
	l.Content = s.Content
	l.CreatedAt = s.CreatedAt
	l.ModifiedAt = s.ModifiedAt
	return l
}

// User represent the structure of the frontend-facing user. Not the stored one.
// It offers an `FromStorage` function to convert itself from the storage payload.
type User struct {
	ID string `json:"id" mapstructure:"id"`
	// The username is the login of the user.
	Username string `json:"username" mapstructure:"username"`
	// Is 2FA enabled on this account.
	// It will be used to determine if it requires another step during the login process
	TwoFactorEnabled bool      `json:"two_factor_enabled" mapstructure:"two_factor_enabled"`
	CreatedAt        time.Time `json:"created_at" mapstructure:"created_at"`
	ModifiedAt       time.Time `json:"modified_at" mapstructure:"modified_at"`
}

func (u User) FromStorage(s models.User) User {
	u.ID = s.ID
	u.Username = s.Username
	u.TwoFactorEnabled = s.TwoFactorEnabled
	u.CreatedAt = s.CreatedAt
	u.ModifiedAt = s.ModifiedAt
	return u
}
