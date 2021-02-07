package service

import (
	"gitlab.com/ziby90/todo-app/model"
	"gitlab.com/ziby90/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type TodoList interface {
	Create(userId uint, list model.TodoList) (uint, error)
	GetAll(userId uint) ([]model.TodoList, error)
	GetById(userId, listId uint) (model.TodoList, error)
	Delete(userId, listId uint) error
	Update(userId, listId uint, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId uint, input model.TodoItem) (uint, error)
	GetAllItems(userId, listId uint) ([]model.TodoItem, error)
	GetById(userId, listId, itemId uint) (*model.TodoItem, error)
	Delete(userId, listId, itemId uint) error
	Update(userId, listId, itemId uint, input model.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
