package model

import (
	"errors"
	"time"
)

type TodoList struct {
	Id          uint    `json:"id" db:"id"`
	Title       string  `json:"title" db:"title" binding:"required"`
	Description *string `json:"description" db:"description"`
	UserId      uint    `json:"user_id" db:"user_id"`
	Created     time.Time
}

type TodoItem struct {
	Id          uint    `json:"id" db:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Done        bool    `json:"done"`
	ListId      uint    `json:"list_id" db:"list_id"`
	Created     time.Time
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("нет параметров для обновления")
	}
	return nil
}
func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("нет параметров для обновления")
	}
	return nil
}
