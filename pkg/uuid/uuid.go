package uuid

import "github.com/satori/go.uuid"

// New returns a new UUID
func New() string {
	return uuid.Must(uuid.NewV4()).String()
}
