package psql

import (
	"log"

	"github.com/andream16/curve-challenge/pkg/psql"
)

var Tables = map[string]string{
	`Cards`: `CREATE TABLE IF NOT EXISTS payment_accounts (
		ID UUID PRIMARY KEY, 
		owner UUID REFERENCES users (ID),
		name text,
		available_balance double precision,
		marked_balance double precision
	)`,
	`Users`: `CREATE TABLE IF NOT EXISTS users (
		ID UUID PRIMARY KEY
	)`,
	`Merchants`: `CREATE TABLE IF NOT EXISTS merchants (
		ID UUID PRIMARY KEY,
		name text,
		location text,
		balance double precision
	)`,
	`Transactions`: `CREATE TABLE IF NOT EXISTS transactions (
		ID UUID PRIMARY KEY, 
		sender UUID REFERENCES users (ID),
		receiver UUID REFERENCES users (ID),
		date text,
		amount double precision,
		type text
	)`,
	`Authorizations`: `CREATE TABLE IF NOT EXISTS authorizations (
		ID UUID PRIMARY KEY, 
		transaction text REFERENCES transactions(ID),
		amount double precision,
		captured double precision,
		catched boolean
	)`,
}

// CreateTables creates Tables mandatory for the program
func CreateTables(svc *psql.PSQL) error {

	for k, v := range Tables {

		log.Printf("Creating table %v ...", k)

		err := svc.CreateTable(v)
		if err != nil {
			return err
		}

		log.Printf("Successfully created table %v !", k)

	}

	return nil

}
