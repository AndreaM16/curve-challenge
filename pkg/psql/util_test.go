package psql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/internal/configuration"
)

func TestNewSessionString(t *testing.T) {

	host := "some_host"
	user := "some_user"
	dbName := "some_db_name"
	sslMode := "some_ssl_mode"
	driverName := "some_driver_name"

	cfg := new(configuration.Configuration)
	cfg.PSQL.Host = host
	cfg.PSQL.User = user
	cfg.PSQL.DBName = dbName
	cfg.PSQL.SSLMode = sslMode
	cfg.PSQL.DriverName = driverName

	out, err := newSessionString(cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, out)
	assert.Contains(t, out, host, user, dbName, sslMode, driverName)

}
