package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {

	owner := "someOwner"
	name := "someName"

	out := NewCard(owner, name)

	assert.NotEmpty(t, out)
	assert.Equal(t, owner, out.Owner)
	assert.Equal(t, name, out.Name)

}

func TestCard_SetOwner(t *testing.T) {

	owner := "someOwner"

	out := new(Card)
	out = out.SetOwner(owner)

	assert.Equal(t, owner, out.Owner)

}

func TestCard_SetName(t *testing.T) {

	name := "someName"

	out := new(Card)
	out = out.SetName(name)

	assert.Equal(t, name, out.Name)

}

func TestCard_SetAvailableBalance(t *testing.T) {

	out := new(Card)
	out = out.SetAvailableBalance()

	assert.Equal(t, EmptyBalance, out.AvailableBalance)

}

func TestCard_SetMarkedBalance(t *testing.T) {

	out := new(Card)
	out = out.SetMarkedBalance()

	assert.Equal(t, EmptyBalance, out.MarkedBalance)

}

func TestCard_SetID(t *testing.T) {

	out := new(Card)
	out = out.SetID()

	assert.NotEmpty(t, out.ID)

}

func TestCard_CanDecrement(t *testing.T) {

	out := new(Card)
	out.SetAvailableBalance().IncrementAvailableBalance(10.0)

	assert.True(t, out.CanDecrement(9.0))

}

func TestCard_IncrementAvailableBalance(t *testing.T) {

	amount := 10.0

	out := new(Card)
	out.SetAvailableBalance().IncrementAvailableBalance(amount)

	assert.Equal(t, amount, out.AvailableBalance)

}

func TestCard_IncrementMarkedBalance(t *testing.T) {

	amount := 10.0

	out := new(Card)
	out.SetMarkedBalance().IncrementMarkedBalance(amount)

	assert.Equal(t, amount, out.MarkedBalance)

}

func TestCard_DecrementAvailableBalance(t *testing.T) {

	amount := 10.0

	out := new(Card)
	out.SetAvailableBalance().IncrementAvailableBalance(amount).DecrementAvailableBalance(9.0)

	assert.Equal(t, 1.0, out.AvailableBalance)

}

func TestCard_DecrementMarkedBalance(t *testing.T) {

	amount := 10.0

	out := new(Card)
	out.SetMarkedBalance().IncrementMarkedBalance(amount).DecrementMarkedBalance(9.0)

	assert.Equal(t, 1.0, out.MarkedBalance)

}
