package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/testdata"
)

func TestTopUp(t *testing.T) {

	amount := 10.0
	name := "someMerchant"
	owner := "c9e35256-e831-49c8-8471-164e17a66e29"

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	c := new(model.Card)
	c.SetName(name).SetOwner(owner)

	createdCard, createdCardErr := CreateCard(svc, c)

	assert.NoError(t, createdCardErr)

	topUp := model.TopUp{
		Card:   createdCard.ID,
		Amount: 10,
	}

	updatedCard, updatedCardErr := TopUp(svc, topUp)

	assert.NoError(t, updatedCardErr)
	assert.Equal(t, amount, updatedCard.AvailableBalance)

}
