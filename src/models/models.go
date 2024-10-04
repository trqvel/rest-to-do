package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-"`
	Tasks    []Task `json:"tasks" gorm:"foreignKey:UserID"`
}
