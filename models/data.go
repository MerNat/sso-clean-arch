package models

import (
	"database/sql"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = apmsql.Open("sqlite3", "./sso.db")
	if err != nil {
		panic(err)
	}
}
