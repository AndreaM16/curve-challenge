package model

import (
	internalError "github.com/andream16/curve-challenge/internal/error"
	"github.com/andream16/curve-challenge/pkg/uuid"
)

// Authorization embeds the payment authorization
type Authorization struct {
	// ID is authorization uuid
	ID string `json:"ID"`
	// Transaction is a reference to the original transaction
	Transaction string `json:"ID"`
	// Amount is the authorized amount
	Amount float64 `json:"amount"`
	// Captured is the captured amount
	Captured float64 `json:"captured"`
}

// SetAmount sets authorization's amount
func (t *Authorization) SetAmount(amount float64) *Authorization {
	t.Amount = amount
	return t
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

// SetCaptured sets authorization's transaction
func (t *Authorization) SetCaptured() *Authorization {
	t.Captured = 0.0
	return t
}

// NewTransaction creates a new transaction
func NewAuthorization(tx string, amount float64) (*Authorization, error) {

	if len(tx) == 0 {
		return nil, internalError.Format(TransactionType, ErrEmptyParameter)
	}
	if amount == 0.0 {
		return nil, internalError.Format(Amount, ErrEmptyParameter)
	}

	out := new(Authorization)
	out.SetID().SetAmount(amount).SetTransaction(tx).SetCaptured()

	return out, nil

}
