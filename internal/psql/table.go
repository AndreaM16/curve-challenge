package psql

import (
	"log"

	"github.com/andream16/curve-challenge/pkg/psql"
)

var Tables = map[string]string{
	`Items`: `CREATE TABLE IF NOT EXISTS items (
		ID UUID PRIMARY KEY, 
		price double precision NOT NULL,
		description text
	)`,
	`PaymentAccounts`: `CREATE TABLE IF NOT EXISTS payment_accounts (
		ID UUID PRIMARY KEY, 
		available_balance double precision NOT NULL,
		blocked_balance double precision NOT NULL
	)`,
	`Users`: `CREATE TABLE IF NOT EXISTS users (
		ID UUID PRIMARY KEY, 
		type text,
		payment_account UUID REFERENCES payment_accounts (ID),
		location text
	)`,
	`Transactions`: `CREATE TABLE IF NOT EXISTS transactions (
		ID UUID PRIMARY KEY, 
		sender UUID REFERENCES users (ID),
		receiver UUID REFERENCES users (ID),
		date text
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
