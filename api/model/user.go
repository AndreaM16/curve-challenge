package model

import (
	internalError "github.com/andream16/curve-challenge/internal/error"
	"github.com/andream16/curve-challenge/pkg/uuid"
)

const (
	MerchantType = "merchant"
	ClassicType  = "classic"
)

// User embeds user information
type User struct {
	// ID is an unique identifier for each user
	ID string `json:"ID,omitempty"`
	// Type is the user type. It can be `Merchant` or `Classic`
	Type string `json:"type,omitempty"`
	// PaymentAccount represents user's Payment Account ID
	PaymentAccount string `json:"payment_account,omitempty"`
	// Location represents user's location. This is filled only it type is `Merchant`
	Location string `json:"location,omitempty"`
}

// SetID Sets user's ID
func (user *User) SetID() *User {
	user.ID = uuid.New()
	return user
}

// SetLocation Sets user's location
func (user *User) SetLocation(location string) *User {
	user.Location = location
	return user
}

func (user *User) SetPaymentAccount(paymentAccount string) *User {
	user.PaymentAccount = paymentAccount
	return user
}

func (user *User) SetType(userType string) *User {
	user.Type = userType
	return user
}

// NewUser creates a new User
func NewUser(account, location, userType string) (*User, error) {

	if len(account) == 0 {
		return nil, internalError.Format(AccountID, ErrEmptyParameter)
	}
	if len(userType) == 0 {
		return nil, internalError.Format(UserType, ErrEmptyParameter)
	}
	if userType != ClassicType && userType != MerchantType {
		return nil, internalError.Format(UserType, ErrBadUserType)
	}
	if len(location) == 0 && userType == MerchantType {
		return nil, internalError.Format(Location, ErrEmptyParameter)
	}

	out := new(User)
	out = out.SetID().SetPaymentAccount(account).SetLocation(location).SetType(userType)

	return out, nil

}
