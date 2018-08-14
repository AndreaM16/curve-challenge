package uuid

import (
	"github.com/satori/go.uuid"
)

// New returns a new UUID
func New() string {
	uuidV4 := uuid.NewV4()
	return uuid.Must(uuidV4, nil).String()
}
