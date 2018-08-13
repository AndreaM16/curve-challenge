package middleware

import (
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
)

// CreateAuthorization creates a new authorization
func CreateAuthorization(svc *psql.PSQL, authorization *model.Authorization) error {

	query := `INSERT INTO authorization(ID, transaction, amount, captured) VALUES ($1, $2, $3, $4)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		authorization.ID,
		authorization.Transaction,
		authorization.Amount,
		authorization.Captured,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}