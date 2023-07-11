package repository

import (
	"context"
	"todo/internal/model"
)

type ITodoItemRepository interface {
	Create(ctx context.Context, request *model.TodoItem) (result *[]model.TodoItem, err error)
	GetAll(ctx context.Context) (result *[]model.TodoItem, err error)
	GetByID(ctx context.Context, ID int64) (result *[]model.TodoItem, err error)
	GetByTitle(ctx context.Context, title string) (result *[]model.TodoItem, err error)
	UpdateByID(ctx context.Context, request *model.TodoItem) (err error)
	DeleteByID(ctx context.Context, ID int64) (err error)
}
