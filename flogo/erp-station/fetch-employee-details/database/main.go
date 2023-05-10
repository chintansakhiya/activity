package database

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

var db *sql.DB
var dbURL string
var err error

const (
	POSTGRES = "postgres"
)

func PostgresDBConnection() (*goqu.Database, error) {
	dbURL = "postgres://isight:isight@localhost:5432/isight?sslmode=disable"
	if db == nil {
		db, err = sql.Open(POSTGRES, dbURL)
		if err != nil {
			return nil, err
		}
		return goqu.New(POSTGRES, db), err
	}
	return goqu.New(POSTGRES, db), err
}
