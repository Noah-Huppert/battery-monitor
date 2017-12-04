package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config holds the settings for the application
type Config struct {
	Db DbConfig
}

// Verify ensures all config values are valid, it returns an error if one is
// not, and nil on success
func (c Config) Verify() error {
	// Check DbConfig
	if err := c.Db.Verify(); err != nil {
		return fmt.Errorf("error verifying db config: %s", err.Error())
	}

	// All good
	return nil
}

// DbConfig holds database connection specific configuration
type DbConfig struct {
	// Host specifies the network address the database can be reached at
	Host string

	// User specifies the user used to authenticate
	User string

	// Password specifies the password used to authenticate
	Password string

	// Db specifies the database to operate on
	Db string
}

// Verify ensures all db config values are valid. It returns an error if one
// is not, nil on success.
//
// For a value to be valid, it must not be empty
func (c DbConfig) Verify() error {
	// Holds empty config fields
	empty := []string{}

	// Host
	if len(c.Host) == 0 {
		empty = append(empty, "Host")
	}

	// User
	if len(c.User) == 0 {
		empty = append(empty, "User")
	}

	// Password
	if len(c.Password) == 0 {
		empty = append(empty, "Password")
	}

	// Db
	if len(c.Db) == 0 {
		empty = append(empty, "Db")
	}

	// If any empty
	if len(empty) > 0 {
		return fmt.Errorf("the db config fields: %s, were empty",
			strings.Join(empty, ","))
	}

	// All good
	return nil
}

// PathEnvKey is the name of the environment variable which can be used to
// modify which file LoadConfig loads.
const PathEnvKey string = "APP_CONFIG_PATH"

// LoadConfig reads a toml configuration file and loads the values in a Config
// struct, which is then returned. If the process is successful a nil error
// will be returned. The error otherwise.
//
// By default it will load from the "config.toml" file in the working
// directory. This can be modified by setting the APP_CONFIG_PATH environment
// variable.
func LoadConfig() (Config, error) {
	var config Config
	var filePath string

	// Default file location
	if cwd, err := os.Getwd(); err == nil {
		filePath = filepath.Join(cwd, "config.toml")
	} else {
		return config, fmt.Errorf("error getting working directory: %s", err.Error())
	}

	// Check if PathEnvKey is set
	if val, set := os.LookupEnv(PathEnvKey); set {
		// Override filePath with custom user provided value
		filePath = val
	}

	// Check if filePath exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return config, fmt.Errorf("config path does not exist: %s", filePath)
	}

	// Read file contents
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading config file \"%s\": %s", filePath, err.Error())
	}

	// Load toml
	if _, err := toml.Decode(string(data), &config); err != nil {
		return config, fmt.Errorf("failed to parse toml config file \"%s\": %s", filePath, err.Error())
	}

	// Verify
	if err := config.Verify(); err != nil {
		return config, fmt.Errof("error verifying config: %s", err.Error())
	}

	// All done, return
	return config, nil
}
