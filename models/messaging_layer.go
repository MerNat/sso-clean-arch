package models

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type messagingRepository struct {
	rMqBroker *amqp.Connection
	broker    *Broker
}

func NewMessagingRepository() MessagingRepo {
	mq := &messagingRepository{
		rMqBroker: RMqBroker,
		broker:    InitBroker,
	}
	//Start to listen incoming msg and send to connected clients
	go mq.ListenAndSend()
	return mq
}

func (mRepo *messagingRepository) ListenAndSend() {
	//We should listen and of course trigger to send to all connected clients once a message arrives from our r.mq
	ch, err := mRepo.rMqBroker.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}
	defer ch.Close()

	//We need to make sure exchange exists
	err = ch.ExchangeDeclare(
		"crypto-info", // name of exchange
		"fanout",      // type of exchange
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		panic("Failed to declare an exchange")
	}

	q, err := ch.QueueDeclare(
		"",    // name (temporary queue)
		false, // durable (we don't need it after server restarts or connection of channel closes)
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic("Failed to declare a queue")
	}

	//Bind exchange with queue
	err = ch.QueueBind(
		q.Name,        // queue name
		"",            // routing key
		"crypto-info", // exchange
		false,
		nil,
	)

	if err != nil {
		panic("Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic("Failed to register a consumer")
	}

	var close chan struct{}

	go func() {
		for d := range msgs {
			mRepo.broker.Notifier <- d.Body
			log.Printf("Sending cryto info to client(s)")
		}
	}()

	<-close
}
