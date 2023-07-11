package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")
	v1.GET("/")
}
