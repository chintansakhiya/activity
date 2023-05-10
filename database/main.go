package database

import (
	"database/sql"
	"strconv"

	"github.com/chintansakhiya/activity/flogo/erp-station/config"
	"github.com/doug-martin/goqu/v9"
)

var db *sql.DB
var dbURL string
var err error

const (
	POSTGRES = "postgres"
)

func PostgresDBConnection(cfg config.DBConfig) (*goqu.Database, error) {
	dbURL = "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + "/" + cfg.Db + "?" + cfg.QueryString
	if db == nil {
		db, err = sql.Open(POSTGRES, dbURL)
		if err != nil {
			return nil, err
		}
		return goqu.New(POSTGRES, db), err
	}
	return goqu.New(POSTGRES, db), err
}
