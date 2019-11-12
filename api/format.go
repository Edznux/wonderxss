package api

import (
	"time"

	"github.com/edznux/wonder-xss/storage/models"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Payload represent the structure of the frontend-facing payload. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Payload struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Hash       string    `json:"hash"` // Used for SRI (sub ressource integrity)
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (p Payload) fromStorage(s models.Payload) Payload {
	p.ID = s.ID
	p.Name = s.Name
	p.Content = s.Content
	p.Hash = s.Hash
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

// Loot represent the structure of the frontend-facing Loots. Not the stored one.
// It offers an `fromStorage` function to convert itself from the storage payload.
type Loot struct {
	ID          string    `json:"id"`
	PayloadID   string    `json:"payload_id"`
	AliasID     string    `json:"alias_id"`
	TriggeredAt time.Time `json:"triggered_at"`
}

func (l Loot) fromStorage(s models.Loot) Loot {
	l.ID = s.ID
	l.PayloadID = s.PayloadID
	l.AliasID = s.AliasID
	l.TriggeredAt = s.TriggeredAt
	return l
}
