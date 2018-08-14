package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

const CATCHED = false

// Authorization embeds the payment authorization
type Authorization struct {
	// ID is authorization uuid
	ID string `json:"ID"`
	// Transaction is a reference to the original transaction that generated the authorization
	Transaction string `json:"ID"`
	// Amount is the authorized amount
	Amount float64 `json:"amount"`
	// Captured is the captured amount
	Captured float64 `json:"captured"`
	// Catched is true if the merchant captured all the amount
	Catched bool `json:"catched"`
}

// SetID sets authorization's ID
func (t *Authorization) SetID() *Authorization {
	t.ID = uuid.New()
	return t
}

// SetTransaction sets authorization's transaction
func (t *Authorization) SetTransaction(tx string) *Authorization {
	t.Transaction = tx
	return t
}

// SetAmount sets authorization's amount
func (t *Authorization) SetAmount(amount float64) *Authorization {
	t.Amount = amount
	return t
}

// SetCaptured sets authorization's transaction
func (t *Authorization) SetCaptured() *Authorization {
	t.Captured = 0.0
	return t
}

// SetCatched sets authorization's catched field
func (t *Authorization) SetCatched() *Authorization {
	t.Catched = CATCHED
	return t
}

// NewTransaction creates a new transaction
func NewAuthorization(tx string, amount float64) *Authorization {

	out := new(Authorization)
	out.SetID().SetTransaction(tx).SetAmount(amount).SetCaptured().SetCatched()

	return out

}
