package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/ziby90/todo-app/model"
	"strings"
)

type TodoItemSql struct {
	db *sqlx.DB
}

func NewTodoItemSql(db *sqlx.DB) *TodoItemSql {
	return &TodoItemSql{db: db}
}

func (r *TodoItemSql) Create(listId uint, input model.TodoItem) (uint, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var idItem uint
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, list_id) VALUES ($1,$2, $3) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description, listId)
	if err := row.Scan(&idItem); err != nil {
		tx.Rollback()
		return 0, err
	}
	return idItem, tx.Commit()
}

func (r *TodoItemSql) GetAllItems(listId uint) ([]model.TodoItem, error) {
	var items []model.TodoItem
	query := fmt.Sprintf("SELECT items.* FROM %s items WHERE items.list_id= $1",
		todoItemsTable)
	err := r.db.Select(&items, query, listId)

	return items, err
}
func (r *TodoItemSql) GetById(listId, itemId uint) (*model.TodoItem, error) {
	var item model.TodoItem
	query := fmt.Sprintf("SELECT * FROM %s  WHERE list_id= $1 AND id=$2",
		todoItemsTable)
	err := r.db.Get(&item, query, listId, itemId)

	return &item, err
}
func (r *TodoItemSql) Delete(listId, itemId uint) error {
	query := fmt.Sprintf("DELETE FROM %s ti WHERE ti.list_id=$1 AND ti.id=$2",
		todoItemsTable)
	_, err := r.db.Exec(query, listId, itemId)

	return err
}
func (r *TodoItemSql) Update(userId, listId, itemId uint, input model.UpdateItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s WHERE ti.list_id=$%d AND ti.id=$%d",
		todoItemsTable, setQuery, argId, argId+1)
	args = append(args, listId, itemId)
	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)
	_, err := r.db.Exec(query, args...)

	return err
}
