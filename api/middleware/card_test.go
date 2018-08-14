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

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	usr, usrErr := CreateUser(svc)

	assert.NoError(t, usrErr)

	c := new(model.Card)
	c.SetName(name).SetOwner(usr.ID)

	out, err := CreateCard(svc, c)

	assert.NoError(t, err)
	assert.Equal(t, name, out.Name)
	assert.Equal(t, usr.ID, out.Owner)

}

func TestGetCard(t *testing.T) {

	name := "someMerchant"

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	usr, usrErr := CreateUser(svc)

	assert.NoError(t, usrErr)

	c := new(model.Card)
	c.SetName(name).SetOwner(usr.ID)

	createdCard, createdCardErr := CreateCard(svc, c)

	assert.NoError(t, createdCardErr)

	card, cardErr := GetCard(svc, createdCard.ID)

	assert.NoError(t, cardErr)
	assert.NotEmpty(t, card)

}

func TestUpdateCard(t *testing.T) {

	amount := 10.0
	name := "someMerchant"

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	usr, usrErr := CreateUser(svc)

	assert.NoError(t, usrErr)

	c := new(model.Card)
	c.SetName(name).SetOwner(usr.ID)

	createdCard, createdCardErr := CreateCard(svc, c)

	assert.NoError(t, createdCardErr)

	createdCard.IncrementAvailableBalance(amount)

	updatedCard, updatedCardErr := UpdateCard(svc, createdCard)

	assert.NoError(t, updatedCardErr)
	assert.Equal(t, amount, updatedCard.AvailableBalance)

}
