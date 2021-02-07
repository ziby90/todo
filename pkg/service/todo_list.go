package service

import (
	"gitlab.com/ziby90/todo-app/model"
	"gitlab.com/ziby90/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId uint, list model.TodoList) (uint, error) {
	return s.repo.Create(userId, list)
}
func (s *TodoListService) GetAll(userId uint) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListService) GetById(userId, listId uint) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
func (s *TodoListService) Delete(userId, listId uint) error {
	return s.repo.Delete(userId, listId)
}
func (s *TodoListService) Update(userId, listId uint, input model.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
