package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/testdata"
)

func TestCreateUser(t *testing.T) {

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	user, userErr := model.NewUser(
		"some_account",
		"some_location",
		"merchant",
	)

	assert.NoError(t, userErr)

	err := CreateUser(svc, user)

	assert.NoError(t, err)

}
