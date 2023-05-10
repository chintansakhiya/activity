package database

import (
	"database/sql"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

var db *sql.DB
var dbURL string
var err error

const (
	POSTGRES    = "postgres"
	Username    = "isight"
	Password    = "isight"
	Host        = "localhost"
	Port        = 5432
	Db          = "isight"
	QueryString = "="
)

func PostgresDBConnection() (*goqu.Database, error) {
	dbURL = "postgres://" + Username + ":" + Password + "@" + Host + ":" + strconv.Itoa(Port) + "/" + Db + "?" + QueryString
	if db == nil {
		db, err = sql.Open(POSTGRES, dbURL)
		if err != nil {
			return nil, err
		}
		return goqu.New(POSTGRES, db), err
	}
	return goqu.New(POSTGRES, db), err
}
