package models

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"github.com/mernat/sso-clean-arch/utils"
)

type repo struct {
	db *sql.DB
}

func NewSQLiteRepository() Repository {
	return &repo{
		db: Db,
	}
}

func (r *repo) CreateUser(user *User) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO users(name, email, password) values(?,?,?)")
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

// func (user *User) CreateUser() (err error) {
// 	stmt, err := Db.Prepare("INSERT INTO users(name, email, password) values(?,?,?)")
// 	if err != nil {
// 		return
// 	}
// 	_, err = stmt.Exec(user.Name, user.Email, utils.Encrypt(user.Password))

// 	if err != nil {
// 		if sqliteErr, ok := err.(sqlite3.Error); ok {
// 			if sqliteErr.Code == sqlite3.ErrConstraint {
// 				err = errors.New("email already exist")
// 			}
// 		}
// 	}
// 	return
// }

func (r *repo) GetUser(user *User) (err error) {
	err = r.db.QueryRow("select name, password from users where email=$1", user.Email).Scan(&user.Name, &user.Password)
	if err != nil {
		err = errors.New("email Not Found")
	}
	return
}

// func (user *User) GetUser() (err error) {
// 	err = Db.QueryRow("select name, password from users where email=$1", user.Email).Scan(&user.Name, &user.Password)
// 	if err != nil {
// 		err = errors.New("email Not Found")
// 	}
// 	return
// }
