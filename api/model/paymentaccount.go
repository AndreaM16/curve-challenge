package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

const EmptyBalance = 0.0

// PaymentAccount embeds a payment account information
type PaymentAccount struct {
	// ID is the unique identifier for a a payment account
	ID string `json:"ID,omitempty"`
	// MarkedBalance is the marked balance. It cannot be used.
	MarkedBalance float64 `json:"marked_balance,omitempty"`
	// AvailableBalance is the available balance
	AvailableBalance float64 `json:"available_balance,omitempty"`
}

// SetMarkedBalance sets a marked balance to the passed one
func (p *PaymentAccount) SetMarkedBalance(balance float64) *PaymentAccount {
	p.MarkedBalance = balance
	return p
}

// SetAvailableBalance sets a available balance to the passed one
func (p *PaymentAccount) SetAvailableBalance(balance float64) *PaymentAccount {
	p.AvailableBalance = balance
	return p
}

// SetID sets account's ID
func (p *PaymentAccount) SetID() *PaymentAccount {
	p.ID = uuid.New()
	return p
}

// NewPaymentAccount creates a new PaymentAccount
func NewPaymentAccount() *PaymentAccount {

	out := new(PaymentAccount)
	out = out.SetID().SetAvailableBalance(EmptyBalance).SetAvailableBalance(EmptyBalance)

	return out

}
