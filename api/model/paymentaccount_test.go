package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPaymentAccount(t *testing.T) {

	out := NewPaymentAccount()

	assert.NotEmpty(t, out)

}

func TestPaymentAccount_SetAvailableBalance(t *testing.T) {

	amount := 10.0

	out := new(PaymentAccount)
	out = out.SetAvailableBalance(amount)

	assert.Equal(t, amount, out.AvailableBalance)

}

func TestPaymentAccount_SetMarkedBalance(t *testing.T) {

	amount := 10.0

	out := new(PaymentAccount)
	out = out.SetMarkedBalance(amount)

	assert.Equal(t, amount, out.MarkedBalance)

}

func TestUser_SetID(t *testing.T) {

	out := new(PaymentAccount)
	out = out.SetID()

	assert.NotEmpty(t, out.ID)

}
