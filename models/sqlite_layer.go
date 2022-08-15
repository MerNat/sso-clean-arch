package models

import (
	"context"
	"database/sql"
	"errors"

	s3 "github.com/mattn/go-sqlite3"
	"github.com/mernat/sso-clean-arch/utils"
	"go.elastic.co/apm"
)

type repo struct {
	db *sql.DB
}

func NewSQLiteRepository() Repository {
	return &repo{
		db: Db,
	}
}

func (r *repo) CreateUser(ctx context.Context, user *User) (err error) {
	span, ctx := apm.StartSpan(ctx, "CreatingUser", "custom")
	defer span.End()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO users(name, email, password) values(?,?,?)")
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, user.Name, user.Email, utils.Encrypt(user.Password))

	if err != nil {
		if sqliteErr, ok := err.(s3.Error); ok {
			if sqliteErr.Code == s3.ErrConstraint {
				err = errors.New("email already exist")
			}
		}
		return
	}
	return tx.Commit()
}

func (r *repo) GetUser(ctx context.Context, user *User) (err error) {
	span, ctx := apm.StartSpan(ctx, "GetUser", "custom")
	defer span.End()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	err = tx.QueryRowContext(ctx, "select name, password from users where email=$1", user.Email).Scan(&user.Name, &user.Password)
	if err != nil {
		err = errors.New("email Not Found")
		return
	}
	return tx.Commit()
}
