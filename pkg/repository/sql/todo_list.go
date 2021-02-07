package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/ziby90/todo-app/model"
	"strings"
)

type TodoListSql struct {
	db *sqlx.DB
}

func NewTodoListSql(db *sqlx.DB) *TodoListSql {
	return &TodoListSql{db: db}
}

func (r *TodoListSql) Create(userId uint, list model.TodoList) (uint, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var idList uint
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, user_id) VALUES ($1,$2,$3) RETURNING id", todoListTables)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, userId)
	if err := row.Scan(&idList); err != nil {
		tx.Rollback()
		return 0, err
	}

	return idList, tx.Commit()
}

func (r *TodoListSql) GetAll(userId uint) ([]model.TodoList, error) {
	var lists []model.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl WHERE tl.user_id = $1",
		todoListTables)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
func (r *TodoListSql) GetById(userId, listId uint) (model.TodoList, error) {
	var list model.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl WHERE tl.user_id= $1 AND tl.id=$2", todoListTables)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
func (r *TodoListSql) Delete(userId, listId uint) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.user_id=$1 AND tl.list_id=$2", todoListTables)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
func (r *TodoListSql) Update(userId, listId uint, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d AND tl.user_id=$%d",
		todoListTables, setQuery, argId, argId+1)
	args = append(args, listId, userId)
	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)
	_, err := r.db.Exec(query, args...)

	return err
}
