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
	Name string `json:"name,omitempty"`
	// MarkedBalance is the marked balance. It cannot be used.
	MarkedBalance float64 `json:"marked_balance"`
	// AvailableBalance is the available balance
	AvailableBalance float64 `json:"available_balance"`
}

// SetID sets card's ID
func (c *Card) SetID() *Card {
	c.ID = uuid.New()
	return c
}

// SetOwner sets card's owner
func (c *Card) SetOwner(owner string) *Card {
	c.Owner = owner
	return c
}

// SetName sets card's name
func (c *Card) SetName(name string) *Card {
	c.Name = name
	return c
}

// SetMarkedBalance sets card's marked balance
func (c *Card) SetMarkedBalance() *Card {
	c.MarkedBalance = EmptyBalance
	return c
}

// SetAvailableBalance sets card's available balance
func (c *Card) SetAvailableBalance() *Card {
	c.AvailableBalance = EmptyBalance
	return c
}

// DecrementAvailableBalance decrements an account available balance
func (c *Card) DecrementAvailableBalance(amount float64) *Card {
	c.AvailableBalance = c.AvailableBalance - amount
	return c
}

// IncrementAvailableBalance increments an account available balance
func (c *Card) IncrementAvailableBalance(amount float64) *Card {
	c.AvailableBalance = c.AvailableBalance + amount
	return c
}

// DecrementMarkedBalance decrements an account marked balance
func (c *Card) DecrementMarkedBalance(amount float64) *Card {
	c.MarkedBalance = c.MarkedBalance - amount
	return c
}

// IncrementAvailableBalance increments an account marked balance
func (c *Card) IncrementMarkedBalance(amount float64) *Card {
	c.MarkedBalance = c.MarkedBalance + amount
	return c
}

// CanDecrement returns true if the passed amount is smaller or equal to the available balance
func (c *Card) CanDecrement(amount float64) bool {
	return c.AvailableBalance >= amount
}

// NewCard creates a new card
func NewCard(owner, name string) *Card {

	out := new(Card)
	out = out.SetID().SetOwner(owner).SetName(name).SetAvailableBalance().SetAvailableBalance()

	return out

}
