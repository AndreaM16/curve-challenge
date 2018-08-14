package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateCard creates a new payment account assigned to a given user
func CreateCard(svc *psql.PSQL) (*string, error) {

	Card := model.NewCard()

	query := `INSERT INTO payment_accounts(ID, available_balance, marked_balance) VALUES ($1, $2, $3)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		Card.ID,
		Card.AvailableBalance,
		Card.MarkedBalance,
	)
	if insertError != nil {
		return nil, insertError
	}

	return &Card.ID, nil

}

// GetCard gets a payment account
func GetCard(svc *psql.PSQL, userID string) (*model.Card, error) {

	user, userErr := GetUser(svc, userID)
	if userErr != nil {
		return nil, userErr
	}

	var account model.Card

	query := `SELECT ID,available_balance,marked_balance FROM payment_accounts WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(user.Card).Scan(&account.ID, &account.AvailableBalance, &account.MarkedBalance)
	if selectErr != nil {
		return nil, selectErr
	}

	return &account, nil

}

// UpdateCard updates an account's balances
func UpdateCard(svc *psql.PSQL, account *model.Card) error {

	query := `UPDATE payment_accounts SET available_balance = $1, marked_balance = $2`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(&account.AvailableBalance, &account.MarkedBalance)
	if insertError != nil {
		return insertError
	}

	return nil

}
