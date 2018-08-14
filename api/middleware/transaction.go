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
	// PAYMENT is used for payment actions
	PAYMENT = "PAYMENT"
)

// TopUp adds money to an user's card
func TopUp(svc *psql.PSQL, topUp *model.TopUp) (*model.Card, error) {

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

	_, txErr := newTransaction(svc, topUp.Amount, "", card.Owner, TOPUP)
	if txErr != nil {
		return nil, txErr
	}

	return card, nil

}

// Pay allows a user to send money to a merchant
func Pay(svc *psql.PSQL, payment *model.Payment) error {

	if payment.Amount == 0 {
		return errors.New("You cannot send 0 amount")
	}

	card, cardErr := GetCard(svc, payment.Card)
	if cardErr != nil {
		return cardErr
	}

	if card.Owner != payment.Sender {
		return errors.New(fmt.Sprintf("You cannot access %v card", payment.Card))
	}

	if !card.CanDecrement(payment.Amount) {
		return errors.New(fmt.Sprintf("You haven't enough funds to cover such amount on your %v card", payment.Card))
	}

	tx, txErr := newTransaction(svc, payment.Amount, payment.Sender, payment.Receiver, PAYMENT)
	if txErr != nil {
		return txErr
	}

	auth := model.NewAuthorization(tx.ID, payment.Card, payment.Amount)

	createAuthErr := CreateAuthorization(svc, auth)
	if createAuthErr != nil {
		return createAuthErr
	}

	card.DecrementAvailableBalance(payment.Amount).IncrementMarkedBalance(payment.Amount)

	_, updateErr := UpdateCard(svc, card)
	if updateErr != nil {
		return updateErr
	}

	return nil

}

// Capture allows a merchant to capture an amount from an authorized payment
func Capture(svc *psql.PSQL, capture *model.Capture) error {

	authz, authzErr := GetAuthorization(svc, capture.Authorization)
	if authzErr != nil {
		return authzErr
	}

	if !authz.CanCapture(capture.Amount) {
		return errors.New(fmt.Sprintf("You cannot capture amount %v on authorization %v since the remaining capture amount available is %v", capture.Amount, capture.Authorization, authz.CaptureAmountAvailable()))
	}

	card, cardErr := GetCard(svc, authz.Card)
	if cardErr != nil {
		return cardErr
	}

	card.DecrementMarkedBalance(capture.Amount)

	_, updateCardErr := UpdateCard(svc, card)
	if updateCardErr != nil {
		return updateCardErr
	}

	tx, txErr := getTransaction(svc, authz.Transaction)
	if txErr != nil {
		return txErr
	}

	merchant, merchantErr := GetMerchant(svc, tx.Receiver)
	if merchantErr != nil {
		return merchantErr
	}

	merchant.IncrementBalance(capture.Amount)

	updateMerchantErr := UpdateMerchant(svc, merchant)
	if updateMerchantErr != nil {
		return updateMerchantErr
	}

	authz.Capture(capture.Amount)

	_, updateAuthErr := UpdateAuthorization(svc, authz)
	if updateAuthErr != nil {
		return updateAuthErr
	}

	return nil

}

// newTransaction writes a new transaction
func newTransaction(svc *psql.PSQL, amount float64, sender, receiver, txType string) (*model.Transaction, error) {

	if txType == TOPUP {
		sender = EXTERNAL
	}

	tx := model.NewTransaction(receiver, sender, txType, amount)

	query := `INSERT INTO transactions (ID,receiver,sender,amount,date,type) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
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
		return nil, insertError
	}

	return tx, nil

}

// getTransaction returns a transaction given its ID
func getTransaction(svc *psql.PSQL, ID string) (*model.Transaction, error) {

	query := `SELECT ID,receiver,sender,amount,date,type FROM transactions WHERE ID = $1`

	var tx model.Transaction

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(ID).Scan(&tx.ID, &tx.Receiver, &tx.Sender, &tx.Amount, &tx.Date, &tx.Type)
	if selectErr != nil {
		return nil, selectErr
	}

	return &tx, nil

}
