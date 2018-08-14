package model

// TopUp embeds an amount
type TopUp struct {
	// Amount is top up's amount
	Amount float64 `json:"amount"`
	// Card is top up's card ID
	Card string `json:"card"`
}
