package models

import (
	"log"
)

type brokerRepository struct {
	dbBroker *Broker
}

func NewBrokerRepository() BrokerRepo {

	b := &brokerRepository{
		dbBroker: InitBroker,
	}

	go b.Listen()
	return b
}

func (brokerRepo *brokerRepository) Listen() {
	for {
		select {
		case s := <-brokerRepo.dbBroker.NewClients:
			// A new client has joined
			brokerRepo.dbBroker.Clients[s] = true
			log.Printf("Client added. %d registered clients", len(brokerRepo.dbBroker.Clients))
		case s := <-brokerRepo.dbBroker.ClosingClients:
			// A client has dettached
			// remove them from our clients map
			delete(brokerRepo.dbBroker.Clients, s)
			log.Printf("Removed client. %d registered clients", len(brokerRepo.dbBroker.Clients))
		case event := <-brokerRepo.dbBroker.Notifier:
			// case for getting a new msg
			// Thus send it to all clients
			for clientMessageChan, _ := range brokerRepo.dbBroker.Clients {
				clientMessageChan <- event
			}
		}
	}
}

func (brokerRepo *brokerRepository) GetBroker() *Broker {
	return brokerRepo.dbBroker
}
