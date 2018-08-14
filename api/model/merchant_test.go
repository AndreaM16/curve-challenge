package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewMerchant(t *testing.T) {

	name := "someName"
	location := "someLocation"

	out := NewMerchant(name, location)

	assert.Equal(t, name, out.Name)
	assert.Equal(t, location, out.Location)

}

func TestMerchant_SetID(t *testing.T) {

	out := new(Merchant)
	out.SetID()

	assert.NotEmpty(t, out.ID)

}

func TestMerchant_SetName(t *testing.T) {

	name := "someName"

	out := new(Merchant)
	out.SetName(name)

	assert.Equal(t, name, out.Name)

}

func TestMerchant_SetLocation(t *testing.T) {

	location := "someLocation"

	out := new(Merchant)
	out.SetLocation(location)

	assert.Equal(t, location, out.Location)

}

func TestMerchant_SetBalance(t *testing.T) {

	out := new(Merchant)
	out.SetBalance()

	assert.Equal(t, EmptyBalance, out.Balance)

}
