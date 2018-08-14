package model

// Payment embeds payment information
type Payment struct {
	// Amount payment's amount
	Amount float64 `json:"amount"`
	// Amount payment's sender
	Sender string `json:"sender"`
	// Amount payment's receiver
	Receiver string `json:"receiver"`
	// Card is the selected card ID
	Card string `json:"card"`
}
