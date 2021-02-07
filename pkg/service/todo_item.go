package service

import (
	"gitlab.com/ziby90/todo-app/model"
	"gitlab.com/ziby90/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId uint, input model.TodoItem) (uint, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, input)
}
func (s *TodoItemService) GetAllItems(userId, listId uint) ([]model.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return []model.TodoItem{}, err
	}
	return s.repo.GetAllItems(listId)
}
func (s *TodoItemService) GetById(userId, listId, itemId uint) (*model.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetById(listId, itemId)
}
func (s *TodoItemService) Delete(userId, listId, itemId uint) error {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return err
	}
	return s.repo.Delete(listId, itemId)
}
func (s *TodoItemService) Update(userId, listId, itemId uint, input model.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return err
	}
	return s.repo.Update(userId, listId, itemId, input)
}
