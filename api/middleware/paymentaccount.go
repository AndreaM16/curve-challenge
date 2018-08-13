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
