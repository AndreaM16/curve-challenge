package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

// Card embeds a payment account information
type Card struct {
	// ID is the unique identifier for a a payment account
	ID string `json:"ID,omitempty"`
	// Owner is the card owner
	Owner string `json:"owner,omitempty"`
	// Name is the card name
	Name string `json:"owner,omitempty"`
	// MarkedBalance is the marked balance. It cannot be used.
	MarkedBalance float64 `json:"marked_balance"`
	// AvailableBalance is the available balance
	AvailableBalance float64 `json:"available_balance"`
}

// SetID sets card's ID
func (p *Card) SetID() *Card {
	p.ID = uuid.New()
	return p
}

// SetOwner sets card's owner
func (p *Card) SetOwner(owner string) *Card {
	p.Owner = owner
	return p
}

// SetName sets card's name
func (p *Card) SetName(name string) *Card {
	p.Name = name
	return p
}

// SetMarkedBalance sets card's marked balance
func (p *Card) SetMarkedBalance() *Card {
	p.MarkedBalance = EmptyBalance
	return p
}

// SetAvailableBalance sets card's available balance
func (p *Card) SetAvailableBalance() *Card {
	p.AvailableBalance = EmptyBalance
	return p
}

// NewCard creates a new card
func NewCard(owner, name string) *Card {

	out := new(Card)
	out = out.SetID().SetOwner(owner).SetName(name).SetAvailableBalance().SetAvailableBalance()

	return out

}
