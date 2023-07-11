package controllers

import "github.com/gin-gonic/gin"

type ITodoItemController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}
