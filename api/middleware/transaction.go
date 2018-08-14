package middleware

import (
	"fmt"

	"github.com/go-errors/errors"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

const (

	// EXTERNAL is used to set the a default sender for Top Ups
	EXTERNAL = "c9e35256-e831-49c8-8471-164e17a66e31"
	// TOPUP is used for top up actions
	TOPUP = "TOPUP"
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

	txErr := newTransaction(svc, topUp.Amount, "", card.Owner, TOPUP)
	if txErr != nil {
		return nil, txErr
	}

	return card, nil

}

// newTransaction writes a new transaction
func newTransaction(svc *psql.PSQL, amount float64, sender, receiver, txType string) error {

	if txType == TOPUP {
		sender = EXTERNAL
	}

	tx := model.NewTransaction(receiver, sender, txType, amount)

	query := `INSERT INTO transactions (ID,receiver,sender,amount,date,type) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		tx.ID,
		tx.Receiver,
		tx.Sender,
		tx.Amount,
		tx.Date,
		tx.Type,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}
