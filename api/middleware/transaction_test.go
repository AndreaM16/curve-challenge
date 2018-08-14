package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/pkg/uuid"
	"github.com/andream16/curve-challenge/testdata"
)

func TestTopUp(t *testing.T) {

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

	topUp := model.TopUp{
		Card:   createdCard.ID,
		Amount: 10,
	}

	updatedCard, updatedCardErr := TopUp(svc, &topUp)

	assert.NoError(t, updatedCardErr)
	assert.Equal(t, amount, updatedCard.AvailableBalance)

}

func TestPay(t *testing.T) {

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

	topUp := model.TopUp{
		Card:   createdCard.ID,
		Amount: amount,
	}

	updatedCard, updatedCardErr := TopUp(svc, &topUp)

	assert.NoError(t, updatedCardErr)
	assert.Equal(t, amount, updatedCard.AvailableBalance)

	merchantIn := model.NewMerchant("some_merchant", "somePlace")

	merchant, merchantErr := CreateMerchant(svc, merchantIn)
	assert.NoError(t, merchantErr)

	payment := &model.Payment{
		Sender:   usr.ID,
		Receiver: merchant.ID,
		Amount:   amount,
		Card:     updatedCard.ID,
	}

	out, err := Pay(svc, payment)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)

}

func TestNewTransaction(t *testing.T) {

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	tx, txErr := newTransaction(svc, 10, uuid.New(), uuid.New(), TOPUP)

	assert.NoError(t, txErr)
	assert.NotEmpty(t, tx)

}
