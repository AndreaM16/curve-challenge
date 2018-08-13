package testdata

import (
	"github.com/andream16/curve-challenge/internal/configuration"
)

var MockConfiguration = &configuration.Configuration{
	Environment: "development",
	Server: configuration.Server{
		Host: "localhost",
		Port: "port",
	},
	PSQL: configuration.PSQL{
		DriverName: "postgres",
		DBName: "curve",
		User: "postgres",
		Host: "localhost",
		SSLMode: "disable",
	},
}

