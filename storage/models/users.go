package models

import (
	"time"
)

//User complement the SafeUser model with its password
type User struct {
	// Extent the SafeUser model
	SafeUser
	// This should never be returned to other parts of the application
	Password string
	// ID should be a string: uuid. Non sequential
	ID string
}

type SafeUser struct {
	Username   string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

//GetUser returns a user model safe to return to the frontend
func (u *User) GetUser() SafeUser {
	su := SafeUser{}

	su.Username = u.Username
	su.CreatedAt = u.CreatedAt
	su.ModifiedAt = u.ModifiedAt

	return su
}
