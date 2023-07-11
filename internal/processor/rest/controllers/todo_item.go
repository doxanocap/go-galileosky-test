package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/cns/errs"
	"todo/internal/manager/interfaces"
	"todo/internal/model"
)

type TodoItemController struct {
	manager interfaces.IManager
}

func InitTodoItemController(manager interfaces.IManager) *TodoItemController {
	return &TodoItemController{
		manager: manager,
	}
}

func (tc *TodoItemController) Create(ctx *gin.Context) {
	var requestBody model.TodoItem

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := tc.manager.Service().TodoItem().Create(ctx, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (tc *TodoItemController) GetAll(ctx *gin.Context) {
	result, err := tc.manager.Service().TodoItem().GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (tc *TodoItemController) GetByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := tc.manager.Service().TodoItem().GetByID(ctx, int64(ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (tc *TodoItemController) UpdateByID(ctx *gin.Context) {
	var requestBody model.TodoItem

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := tc.manager.Service().TodoItem().UpdateByID(ctx, &requestBody)
	if err != nil {
		if errs.IsHttpNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func (tc *TodoItemController) DeleteByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = tc.manager.Service().TodoItem().DeleteByID(ctx, int64(ID))
	if err != nil {
		if errs.IsHttpNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
