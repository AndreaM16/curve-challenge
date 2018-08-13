package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {

	account := "some_account"
	location := "some_location"
	userType := "merchant"

	user, err := NewUser(account, location, userType)

	assert.NoError(t, err)
	assert.Equal(t, account, user.PaymentAccount)
	assert.Equal(t, location, user.Location)
	assert.Equal(t, userType, user.Type)
	assert.NotEmpty(t, user.ID)

}
