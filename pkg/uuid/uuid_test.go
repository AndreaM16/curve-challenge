package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	res := New()

	assert.NotEmpty(t, res)

}
