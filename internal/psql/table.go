package psql

import (
	"github.com/andream16/curve-challenge/pkg/psql"
)

var Tables = []string{
	`CREATE TABLE IF NOT EXISTS users (
		ID UUID PRIMARY KEY
	)`,
	`CREATE TABLE IF NOT EXISTS cards (
		ID UUID PRIMARY KEY, 
		owner UUID REFERENCES users (ID),
		name text,
		available_balance double precision,
		marked_balance double precision
	)`,
	`CREATE TABLE IF NOT EXISTS merchants (
		ID UUID PRIMARY KEY,
		name text,
		location text,
		balance double precision
	)`,
	`CREATE TABLE IF NOT EXISTS transactions (
		ID UUID PRIMARY KEY, 
		sender UUID REFERENCES users (ID),
		receiver UUID REFERENCES users (ID),
		date text,
		amount double precision,
		type text
	)`,
	`CREATE TABLE IF NOT EXISTS authorizations (
		ID UUID PRIMARY KEY, 
		transaction UUID REFERENCES transactions (ID),
		amount double precision,
		captured double precision,
		catched boolean
	)`,
}

// CreateTables creates Tables mandatory for the program
func CreateTables(svc *psql.PSQL) error {

	for _, v := range Tables {

		err := svc.CreateTable(v)
		if err != nil {
			return err
		}

	}

	return nil

}
