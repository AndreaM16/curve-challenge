package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

// Authorization embeds the payment authorization
type Authorization struct {
	// ID is authorization uuid
	ID string `json:"ID"`
	// Transaction is a reference to the original transaction that generated the authorization
	Transaction string `json:"transaction"`
	// Amount is the authorized amount
	Amount float64 `json:"amount"`
	// Captured is the captured amount
	Captured float64 `json:"captured"`
	// Catched is true if the merchant captured all the amount
	Catched bool `json:"catched"`
	// Card represents the card where we have to pick money from
	Card string `json:"card"`
}

// SetID sets authorization's ID
func (a *Authorization) SetID() *Authorization {
	a.ID = uuid.New()
	return a
}

// SetTransaction sets authorization's transaction
func (a *Authorization) SetTransaction(tx string) *Authorization {
	a.Transaction = tx
	return a
}

// SetAmount sets authorization's amount
func (a *Authorization) SetAmount(amount float64) *Authorization {
	a.Amount = amount
	return a
}

// SetCaptured sets authorization's transaction
func (a *Authorization) SetCaptured() *Authorization {
	a.Captured = 0.0
	return a
}

// SetCard sets authorization's transaction
func (a *Authorization) SetCard(card string) *Authorization {
	a.Card = card
	return a
}

// SetCatched sets authorization's catched field
func (a *Authorization) SetCatched() *Authorization {
	a.Catched = false
	return a
}

// CaptureAmountAvailable represents the amount that can be captured
func (a *Authorization) CaptureAmountAvailable() float64 {
	return a.Amount - a.Captured
}

// CanCapture returns true if it's possible to perform the capture
func (a *Authorization) CanCapture(amount float64) bool {
	return !a.Catched && (a.CaptureAmountAvailable() >= amount)
}

// Capture captures funds from authorization
func (a *Authorization) Capture(amount float64) *Authorization {
	a.Captured = a.Captured + amount
	if a.Captured == a.Amount {
		a.Catched = true
	}
	return a
}

// CanRefund returns true if it's possible to perform a refund
func (a *Authorization) CanRefund(amount float64) bool {
	return a.Captured >= amount
}

// NewAuthorization creates a new authorization
func NewAuthorization(tx, card string, amount float64) *Authorization {

	out := new(Authorization)
	out.SetID().SetTransaction(tx).SetAmount(amount).SetCaptured().SetCatched().SetCard(card)

	return out

}
