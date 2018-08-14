package model

// Refund represents the information used by a merchant to refund a user
type Refund struct {
	// Amount is the amount to be refunded
	Amount float64 `json:"amount"`
	// Authorization is the authorization ID
	Authorization string `json:"authorization"`
}
