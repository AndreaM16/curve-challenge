package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

// User embeds user information
type User struct {
	// ID is an unique identifier for each user
	ID string `json:"ID,omitempty"`
}

// SetID Sets user's ID
func (user *User) SetID() *User {
	user.ID = uuid.New()
	return user
}

// NewUser creates a new User
func NewUser() *User {

	out := new(User)
	out = out.SetID()

	return out

}
