package model

// Revert is used by the merchant to reverse a part of the initial authorized amount
type Revert struct {
	// Amount the amount to be reversed
	Amount float64 `json:"amount"`
	// Authorization the UUID of the authorization
	Authorization string `json:"authorization"`
}
