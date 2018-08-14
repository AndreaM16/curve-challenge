package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateAuthorization creates a new authorization
func CreateAuthorization(svc *psql.PSQL, authorization *model.Authorization) error {

	query := `INSERT INTO authorizations (ID,transaction,amount,captured,catched,card) VALUES ($1, $2, $3, $4, $5, $6)`

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
		authorization.Card,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}

// GetAuthorization returns an authorization given its ID
func GetAuthorization(svc *psql.PSQL, ID string) (*model.Authorization, error) {

	query := `SELECT ID,transaction,amount,captured,catched,card FROM authorizations WHERE ID = $1`

	var auth model.Authorization

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(ID).Scan(&auth.ID, &auth.Transaction, &auth.Amount, &auth.Captured, &auth.Catched, &auth.Card)
	if selectErr != nil {
		return nil, selectErr
	}

	return &auth, nil

}

// UpdateAuthorization updates an authorization
func UpdateAuthorization(svc *psql.PSQL, authorization *model.Authorization) (*model.Authorization, error) {

	query := `UPDATE authorizations SET captured=$2, catched=$3 WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(&authorization.ID, &authorization.Captured, &authorization.Catched)
	if updateErr != nil {
		return nil, updateErr
	}

	return authorization, nil

}
