package models

import (
	"database/sql"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"
)

var Db *sql.DB
var InitBroker *Broker
var RMqBroker *amqp.Connection

func init() {
	var err error
	//Init broker
	InitBroker = &Broker{
		Notifier:       make(chan []byte, 1),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}

	//Init R.MQ connection
	RMqBroker, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("Failed to connect to RabbitMQ")
	}
	//Init database connection
	Db, err = apmsql.Open("sqlite3", "./sso.db")
	if err != nil {
		panic(err)
	}
}
