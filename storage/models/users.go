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
	// TOTPSecret represent the shared secret for TOTP
	TOTPSecret string
}

type SafeUser struct {
	// ID should be a string: uuid. Non sequential
	ID string `json:"id"`
	// The username is the login of the user.
	Username string `json:"username"`
	// Is 2FA enabled on this account.
	// It will be used to determine if it requires another step during the login process
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
	CreatedAt        time.Time `json:"created_at"`
	ModifiedAt       time.Time `json:"modified_at"`
}

//GetUser returns a user model safe to return to the frontend
func (u *User) GetUser() SafeUser {
	su := SafeUser{}

	su.Username = u.Username
	su.TwoFactorEnabled = u.TwoFactorEnabled
	su.CreatedAt = u.CreatedAt
	su.ModifiedAt = u.ModifiedAt

	return su
}
