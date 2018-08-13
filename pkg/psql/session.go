package psql

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/andream16/curve-challenge/internal/configuration"
)

// PSQL embeds connection to psql database
type PSQL struct {
	*sql.DB
}

// SetDB sets the DB to PSQL
func (svc *PSQL) SetDB(dbSession *sql.DB) *PSQL {
	svc.DB = dbSession
	return svc
}

// New returns a new *PSQL initialized with the configuration
func New(cfg *configuration.Configuration) (*PSQL, error) {

	sessionString, sessionStringErr := newSessionString(cfg)
	if sessionStringErr != nil {
		return nil, sessionStringErr
	}

	db, dbErr := sql.Open(
		cfg.PSQL.DriverName,
		sessionString,
	)
	if dbErr != nil {
		return nil, dbErr
	}

	out := new(PSQL)
	out = out.SetDB(db)

	return out, nil

}
