package middleware

import (
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
)

// CreateTransaction creates a transaction
func CreateTransaction(svc *psql.PSQL, tx *model.Transaction) error {

	newTx, newTxErr := model.NewTransaction(tx.Receiver, tx.Sender, tx.Type, tx.Amount)
	if newTxErr != nil{
		return newTxErr
	}

	query := `INSERT INTO transactions(ID, sender, receiver, amount, date, type) VALUES ($1, $2, $3, %4, $5, $6)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		newTx.ID,
		newTx.Sender,
		newTx.Receiver,
		newTx.Amount,
		newTx.Date,
		newTx.Type,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}
