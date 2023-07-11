package interfaces

import "todo/internal/manager/interfaces/service"

type IService interface {
	TodoItem() service.ITodoItemService
}
