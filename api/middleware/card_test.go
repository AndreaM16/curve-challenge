package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/testdata"
)

func TestCreateCard(t *testing.T) {

	name := "someMerchant"
	owner := "c9e35256-e831-49c8-8471-164e17a66e29"

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
