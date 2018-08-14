package middleware

import (
	"fmt"

	"github.com/go-errors/errors"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// TopUp adds money to an user's card
func TopUp(svc *psql.PSQL, topUp model.TopUp) (*model.Card, error) {

	if topUp.Amount == 0 {
		return nil, errors.New("You cannot top up a 0 amount")
	}

	card, cardErr := GetCard(svc, topUp.Card)
	if cardErr != nil {
		return nil, errors.New(fmt.Sprintf("Card %v does not exist", topUp.Card))
	}

	card.IncrementAvailableBalance(topUp.Amount)

	_, updateErr := UpdateCard(svc, card)
	if updateErr != nil {
		return nil, updateErr
	}

	return card, nil

}
