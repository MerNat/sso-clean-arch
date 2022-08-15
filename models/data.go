package models

import (
	"database/sql"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"
)

var Db *sql.DB
var InitBroker *Broker

func init() {
	var err error
	//Init broker
	InitBroker = &Broker{
		Notifier:       make(chan []byte, 1),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}
	Db, err = apmsql.Open("sqlite3", "./sso.db")
	if err != nil {
		panic(err)
	}
}
