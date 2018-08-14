package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/testdata"
	"github.com/andream16/curve-challenge/pkg/psql"
)

func TestCreateUser(t *testing.T) {

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	out, err := CreateUser(svc)

	assert.NoError(t, err)

	assert.NotEmpty(t, out)

}