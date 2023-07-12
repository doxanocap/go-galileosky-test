package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	todo := v1.Group("/todo")
	{
		todo.POST("", h.ctl.TodoItem().Create)
		todo.GET("/all", h.ctl.TodoItem().GetAll)
		todo.GET("/:id", h.ctl.TodoItem().GetByID)
		todo.PUT("/:id", h.ctl.TodoItem().UpdateByID)
		todo.DELETE("/:id", h.ctl.TodoItem().DeleteByID)
	}
}
