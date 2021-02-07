package model

import "time"

type User struct {
	Id       uint   `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Created  time.Time
}
