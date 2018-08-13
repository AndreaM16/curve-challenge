package model

const (

	// ErrBadUserType is used when use type is not `Classic` or `Merchant`
	ErrBadUserType = "BadUserType"

	// ErrEmptyParameter is used when a required parameter is empty
	ErrEmptyParameter = "EmptyParameter"

	// ErrBadAmount is used when an amount is 0 and therefore not allowed
	ErrBadAmount = "BadAmount"
)
