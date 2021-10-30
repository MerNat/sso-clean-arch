package models

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/mattn/go-sqlite3"
	"github.com/mernat/sso-clean-arch/utils"
)

//UserContextKey declares context
type UserContextKey struct{}

type UserClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (user *User) CreateUser() (err error) {
	stmt, err := Db.Prepare("INSERT INTO users(name, email, password) values(?,?,?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(user.Name, user.Email, utils.Encrypt(user.Password))

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				err = errors.New("email already exist")
			}
		}
	}
	return
}

func (user *User) GetUser() (err error) {
	err = Db.QueryRow("select name, password from users where email=$1", user.Email).Scan(&user.Name, &user.Password)
	if err != nil {
		err = errors.New("email Not Found")
	}
	return
}
