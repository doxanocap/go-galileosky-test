package interfaces

import "todo/internal/manager/interfaces/repository"

type IRepository interface {
	TodoItem() repository.ITodoItemRepository
}
