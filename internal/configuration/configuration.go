package configuration

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/tkanos/gonfig"
)

const confFileName = "configuration.json"

// Configuration embeds environment variables used to initialize the system
type Configuration struct {
	// Environment represents the working environment
	Environment string `json:"environment"`
	// Server embeds environment variables used to initialize the server
	Server Server `json:"SERVER"`
	// PSQL embeds environment variables used to initialize PSQL connection
	PSQL PSQL `json:"PSQL"`
}

// Server embeds server information
type Server struct {
	// Host is the host of the server
	Host string `json:"HOST"`
	// Port is the port where the server should listen
	Port string `json:"PORT"`
}

// PSQL embeds PSQL information
type PSQL struct {
	// DriverName represents PSQL driver name
	DriverName string `json:"DRIVERNAME"`
	// DBName represents PSQL database table name
	DBName string `json:"DBNAME"`
	// User represents PSQL User
	User string `json:"USER"`
	// Host represents PSQL database host
	Host string `json:"HOST"`
	// SSLMode has value `disable` if no authentication is required
	SSLMode string `json:"SSLMODE"`
}

// Get returns a *Configuration filled with values fetched from a configuration file confFileName
func Get() (*Configuration, error) {

	var cfg Configuration

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), confFileName)

	if err := gonfig.GetConf(filePath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
