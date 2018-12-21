package configuration

import "os"

// DB is the configuration for the datastore
var DB DBConfig

func init() {
	DB = GetDBConfig()
}

// DBConfig is the configuration struct for the datastore
type DBConfig struct {
	Adapter    string
	Host       string
	User       string
	DBName     string
	DBPassword string
	SslMode    string
}

// GetDBConfig returns configuration for the datastore
func GetDBConfig() DBConfig {
	c := DBConfig{
		Adapter:    os.Getenv("DBTYPE"),
		Host:       os.Getenv("DBHOST"),
		User:       os.Getenv("DBUSER"),
		DBName:     os.Getenv("DBNAME"),
		DBPassword: os.Getenv("DBPASSWORD"),
		SslMode:    os.Getenv("DBSSL"),
	}

	// set defaults
	if c.Adapter == "" {
		c.Adapter = "postgres"
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.User == "" {
		c.User = "postgres"
	}
	if c.DBName == "" {
		c.DBName = "authenz_dev"
	}
	if c.DBPassword == "" {
		c.DBPassword = ""
	}
	if c.SslMode == "" {
		c.SslMode = "disable"
	}

	return c
}
