package model

// Capture embeds information used by the merchant to perform a capture action
type Capture struct {
	// Merchant represents merchant's ID
	Merchant string `json:"merchant"`
	// Authorization represents authorization's ID
	Authorization string `json:"authorization"`
	// Amount represents the amount to be captured
	Amount float64 `json:"amount"`
}
