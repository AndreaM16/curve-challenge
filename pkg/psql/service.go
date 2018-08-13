package psql

import (
	internalErr "github.com/andream16/curve-challenge/internal/error"
)

// CreateTable creates a table given a query containing the SQL syntax to create a table.
// E.g. : `CREATE TABLE example (someKey text PRIMARY KEY, someField text)`
func (svc *PSQL) CreateTable(query string) error {

	if len(query) == 0 {
		return internalErr.Format(ErrEmptyParameter, Query)
	}

	_, err := svc.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
