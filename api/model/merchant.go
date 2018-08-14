package model

import (
	"github.com/andream16/curve-challenge/pkg/uuid"
)

// Merchant embeds merchant information
type Merchant struct {
	// ID is an unique identifier for each merchant
	ID string `json:"ID,omitempty"`
	// Name represents merchant's name
	Name string `json:"name,omitempty"`
	// Location represents merchant's location
	Location string `json:"location,omitempty"`
	// Balance represents merchant's balance
	Balance float64 `json:"balance,omitempty"`
}

// SetID Sets merchant's ID
func (merchant *Merchant) SetID() *Merchant {
	merchant.ID = uuid.New()
	return merchant
}

// SetName Sets merchant's name
func (merchant *Merchant) SetName(name string) *Merchant {
	merchant.Name = name
	return merchant
}

// SetLocation Sets merchant's location
func (merchant *Merchant) SetLocation(location string) *Merchant {
	merchant.Location = location
	return merchant
}

// SetLocation Sets merchant's location
func (merchant *Merchant) SetBalance() *Merchant {
	merchant.Balance = EmptyBalance
	return merchant
}

// NewMerchant creates a new merchant
func NewMerchant(name, location string) *Merchant {

	out := new(Merchant)
	out = out.SetID().SetName(name).SetLocation(location).SetBalance()

	return out

}
