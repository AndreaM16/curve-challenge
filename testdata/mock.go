package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/internal/configuration"
)

func MockConfiguration(t *testing.T) *configuration.Configuration {

	t.Helper()

	cfg, err := configuration.Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)

	return cfg

}
