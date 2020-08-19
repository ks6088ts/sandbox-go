package gql

import (
	"database/sql"
	"fmt"
)

// DatabaseConfig ...
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func getDatabase(config DatabaseConfig) (*sql.DB, error) {
	return sql.Open(config.Driver,
		fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Dbname))
}
