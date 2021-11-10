package models

import "github.com/mernat/sso-clean-arch/utils"

type IRepo struct {
	m map[int]*User
}

func NewInmemRepository() *IRepo {
	var m = map[int]*User{}
	return &IRepo{
		m: m,
	}
}

func (r *IRepo) CreateUser(user *User) (err error) {
	user.Password = utils.Encrypt(user.Password)
	r.m[len(r.m)+1] = user
	return nil
}

func (r *IRepo) GetUser(user *User) (err error) {
	for _, v := range r.m {
		if v.Email == user.Email {
			user.Name = v.Name
			user.Password = v.Password
		}
	}
	return nil
}
