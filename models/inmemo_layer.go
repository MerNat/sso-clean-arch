package models

import (
	"context"

	"github.com/mernat/sso-clean-arch/utils"
)

type IRepo struct {
	m map[int]*User
}

func NewInmemRepository() *IRepo {
	var m = map[int]*User{}
	return &IRepo{
		m: m,
	}
}

func (r *IRepo) CreateUser(ctx context.Context, user *User) (err error) {
	user.Password = utils.Encrypt(user.Password)
	r.m[len(r.m)+1] = user
	return nil
}

func (r *IRepo) GetUser(ctx context.Context, user *User) (err error) {
	for _, v := range r.m {
		if v.Email == user.Email {
			user.Name = v.Name
			user.Password = v.Password
		}
	}
	return nil
}
