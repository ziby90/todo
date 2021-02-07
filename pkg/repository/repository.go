package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/ziby90/todo-app/model"
	"gitlab.com/ziby90/todo-app/pkg/repository/sql"
)

type Authorization interface {
	CreateUser(user model.User) (uint, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userId uint, list model.TodoList) (uint, error)
	GetAll(userId uint) ([]model.TodoList, error)
	GetById(userId, listId uint) (model.TodoList, error)
	Delete(userId, listId uint) error
	Update(userId, listId uint, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId uint, input model.TodoItem) (uint, error)
	GetAllItems(listId uint) ([]model.TodoItem, error)
	GetById(listId, itemId uint) (*model.TodoItem, error)
	Delete(listId, itemId uint) error
	Update(userId, listId, itemId uint, input model.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: sql.NewAuthSql(db),
		TodoList:      sql.NewTodoListSql(db),
		TodoItem:      sql.NewTodoItemSql(db),
	}
}
