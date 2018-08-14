package psql

import (
	"fmt"

	"github.com/andream16/curve-challenge/internal/configuration"
	internalErr "github.com/andream16/curve-challenge/internal/error"
)

// CONNECTIONSTRING represents the default psql string to be filled
const CONNECTIONSTRING = "host=%s user=%s dbname=%s sslmode=%s"

// newSessionString returns a new connection string initialized using the configuration
func newSessionString(cfg *configuration.Configuration) (string, error) {

	if len(cfg.PSQL.Host) == 0 {
		return "", internalErr.Format(ErrEmptyParameter, Host)
	}
	if len(cfg.PSQL.User) == 0 {
		return "", internalErr.Format(ErrEmptyParameter, User)
	}
	if len(cfg.PSQL.DBName) == 0 {
		return "", internalErr.Format(ErrEmptyParameter, User)
	}
	if len(cfg.PSQL.SSLMode) == 0 {
		return "", internalErr.Format(ErrEmptyParameter, User)
	}
	if len(cfg.PSQL.DriverName) == 0 {
		return "", internalErr.Format(ErrEmptyParameter, DriverName)
	}

	return fmt.Sprintf(CONNECTIONSTRING, cfg.PSQL.Host, cfg.PSQL.User, cfg.PSQL.DBName, cfg.PSQL.SSLMode), nil

}
