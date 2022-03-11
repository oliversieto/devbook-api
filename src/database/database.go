package database

import (
	"database/sql"
	"devbook-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dbConnection, err := sql.Open("mysql", config.DataSourceName)

	if err != nil {
		return nil, err
	}

	if err = dbConnection.Ping(); err != nil {
		dbConnection.Close()
		return nil, err
	}

	return dbConnection, nil
}
