package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {

	receiver := "someReceiver"
	sender := "someSender"
	amount := 10.0
	txType := "someTxType"

	out := NewTransaction(
		receiver,
		sender,
		txType,
		amount,
	)

	assert.Equal(t, receiver, out.Receiver)
	assert.Equal(t, sender, out.Sender)
	assert.Equal(t, txType, out.Type)
	assert.Equal(t, amount, out.Amount)

}

func TestTransaction_SetID(t *testing.T) {

	out := new(Transaction)
	out.SetID()

	assert.NotEmpty(t, out.ID)

}

func TestTransaction_SetReceiver(t *testing.T) {

	receiver := "someReceiver"

	out := new(Transaction)
	out.SetReceiver(receiver)

	assert.Equal(t, receiver, out.Receiver)

}

func TestTransaction_SetSender(t *testing.T) {

	sender := "someSender"

	out := new(Transaction)
	out.SetSender(sender)

	assert.Equal(t, sender, out.Sender)

}

func TestTransaction_SetType(t *testing.T) {

	txType := "someTxType"

	out := new(Transaction)
	out.SetType(txType)

	assert.Equal(t, txType, out.Type)

}

func TestTransaction_SetAmount(t *testing.T) {

	amount := 10.0

	out := new(Transaction)
	out.SetAmount(amount)

	assert.Equal(t, amount, out.Amount)

}

func TestTransaction_SetDate(t *testing.T) {

	out := new(Transaction)
	out.SetDate()

	assert.NotEmpty(t, out.Date)

}
