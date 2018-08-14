package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	user := NewUser()

	assert.NotEmpty(t, user.ID)

}

func TestUser_SetID(t *testing.T) {

	out := new(User)
	out = out.SetID()

	assert.NotEmpty(t, out.ID)

}
