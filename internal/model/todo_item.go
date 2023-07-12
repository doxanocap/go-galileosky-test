package model

import "time"

type TodoItem struct {
	Id          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description string     `json:"description" db:"description"`
	Done        bool       `json:"done" db:"done"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type CreateTodoItem struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateTodoItem struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type CustomError struct {
	CustomError string `json:"error"`
}
