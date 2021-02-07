package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ziby90/todo-app/model"
)

type AuthSql struct {
	db *sqlx.DB
}

func NewAuthSql(db *sqlx.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (r *AuthSql) CreateUser(user model.User) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3) RETURNING id", usersTables)

	row := r.db.QueryRow(query, user.Name, user.UserName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthSql) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password_hash=$2", usersTables)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
