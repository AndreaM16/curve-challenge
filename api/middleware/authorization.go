package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateAuthorization creates a new authorization
func CreateAuthorization(svc *psql.PSQL, authorization *model.Authorization) error {

	query := `INSERT INTO authorizations (ID,transaction,amount,captured,catched) VALUES ($1, $2, $3, $4, $5)`

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
		authorization.Catched,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}
