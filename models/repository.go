package models

type Repository interface {
	CreateUser(user *User) (err error)
	GetUser(user *User) (err error)
}