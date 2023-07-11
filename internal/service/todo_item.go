package service

import (
	"context"
	"todo/internal/cns/errs"
	"todo/internal/manager/interfaces"
	"todo/internal/model"
)

type TodoItemService struct {
	manager interfaces.IManager
}

func InitTodoItemService(manager interfaces.IManager) *TodoItemService {
	return &TodoItemService{
		manager: manager,
	}
}

func (ts *TodoItemService) Create(ctx context.Context, request *model.TodoItem) (item *model.TodoItem, err error) {
	result, err := ts.manager.Repository().TodoItem().Create(ctx, request)
	if err != nil {
		return
	}
	item = &(*result)[0]
	return
}

func (ts *TodoItemService) GetAll(ctx context.Context) (items *[]model.TodoItem, err error) {
	return ts.manager.Repository().TodoItem().GetAll(ctx)
}

func (ts *TodoItemService) GetByID(ctx context.Context, ID int64) (item *model.TodoItem, err error) {
	result, err := ts.manager.Repository().TodoItem().GetByID(ctx, ID)
	if result == nil || len(*result) == 0 {
		return nil, errs.HttpNotFound("item")
	}
	item = &(*result)[0]
	return
}

func (ts *TodoItemService) GetByTitle(ctx context.Context, title string) (item *model.TodoItem, err error) {
	result, err := ts.manager.Repository().TodoItem().GetByTitle(ctx, title)
	if result == nil || len(*result) == 0 {
		return nil, errs.HttpNotFound("item")
	}

	item = &(*result)[0]
	return
}

func (ts *TodoItemService) UpdateByID(ctx context.Context, request *model.TodoItem) (err error) {
	result, err := ts.manager.Repository().TodoItem().GetByID(ctx, request.Id)
	if result == nil || len(*result) == 0 {
		return errs.HttpNotFound("item")
	}

	return ts.manager.Repository().TodoItem().UpdateByID(ctx, request)
}

func (ts *TodoItemService) DeleteByID(ctx context.Context, ID int64) (err error) {
	result, err := ts.manager.Repository().TodoItem().GetByID(ctx, ID)
	if result == nil || len(*result) == 0 {
		return errs.HttpNotFound("item")
	}
	return ts.manager.Repository().TodoItem().DeleteByID(ctx, ID)
}
