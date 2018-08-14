package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/testdata"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
)

func TestCreateCard(t *testing.T) {

	name := "someMerchant"
	owner := "someOwner"

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	c := new(model.Card)
	c.SetName(name).SetOwner(owner)

	out, err := CreateCard(svc, c)

	assert.NoError(t, err)
	assert.Equal(t, name, out.Name)
	assert.Equal(t, owner, out.Owner)

}
