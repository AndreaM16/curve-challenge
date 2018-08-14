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
