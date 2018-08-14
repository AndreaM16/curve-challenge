package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthorization(t *testing.T) {

	tx := "someTx"
	amount := 10.0

	out := NewAuthorization(tx, amount)

	assert.NotEmpty(t, out)
	assert.Equal(t, tx, out.Transaction)
	assert.Equal(t, amount, out.Amount)

}

func TestAuthorization_SetID(t *testing.T) {

	out := new(Authorization)
	out = out.SetID()

	assert.NotEmpty(t, out.ID)

}

func TestAuthorization_SetTransaction(t *testing.T) {

	tx := "someTx"

	out := new(Authorization)
	out = out.SetTransaction(tx)

	assert.Equal(t, tx, out.Transaction)

}

func TestAuthorization_SetAmount(t *testing.T) {

	amount := 10.0

	out := new(Authorization)
	out = out.SetAmount(amount)

	assert.Equal(t, amount, out.Amount)

}

func TestAuthorization_SetCatched(t *testing.T) {

	out := new(Authorization)
	out = out.SetCatched()

	assert.Equal(t, false, out.Catched)

}

func TestAuthorization_SetCaptured(t *testing.T) {

	out := new(Authorization)
	out = out.SetCaptured()

	assert.Equal(t, EmptyBalance, out.Captured)

}
