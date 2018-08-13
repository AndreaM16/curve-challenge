package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	res := New()

	assert.NotEmpty(t, res)

}
