package handler

const (

	// InvalidParameters is used when one or more parameters are invalid
	InvalidParameters = "One or more parameters are invalid"

	// MissingParameters is used when one or more parameters are missing
	MissingParameters = "One or more parameters are missing"

	// InsufficientFounds is used when a transaction can't be performed since sender's hasn't enough money
	InsufficientFounds = "Insufficient founds to perform the payment"

	// ForbiddenAction is used when user performs a forbidden action
	ForbiddenAction = "Forbidden action"
)
