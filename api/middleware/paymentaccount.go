package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreatePaymentAccount creates a new payment account assigned to a given user
func CreatePaymentAccount(svc *psql.PSQL) (*string, error) {

	paymentAccount := model.NewPaymentAccount()

	query := `INSERT INTO payment_accounts(ID, available_balance, marked_balance) VALUES ($1, $2, $3)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		paymentAccount.ID,
		paymentAccount.AvailableBalance,
		paymentAccount.MarkedBalance,
	)
	if insertError != nil {
		return nil, insertError
	}

	return &paymentAccount.ID, nil

}

// GetPaymentAccount gets a payment account
func GetPaymentAccount(svc *psql.PSQL, userID string) (*model.PaymentAccount, error) {

	user, userErr := GetUser(svc, userID)
	if userErr != nil {
		return nil, userErr
	}

	var account model.PaymentAccount

	query := `SELECT ID,available_balance,marked_balance FROM payment_accounts WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(user.PaymentAccount).Scan(&account.ID, &account.AvailableBalance, &account.MarkedBalance)
	if selectErr != nil {
		return nil, selectErr
	}

	return &account, nil

}

// UpdatePaymentAccount updates an account's balances
func UpdatePaymentAccount(svc *psql.PSQL, account *model.PaymentAccount) error {

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
