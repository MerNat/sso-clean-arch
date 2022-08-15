package models

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) (err error)
	GetUser(ctx context.Context, user *User) (err error)
}

type BrokerRepo interface {
	Listen()
	GetBroker() *Broker
}
