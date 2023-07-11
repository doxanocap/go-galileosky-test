package manager

import (
	"sync"
	"todo/internal/manager/interfaces"
	"todo/internal/manager/interfaces/service"
	service2 "todo/internal/service"
)

type ServiceManager struct {
	manager interfaces.IManager

	todoItem       service.ITodoItemService
	todoItemRunner sync.Once
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}

func (sm *ServiceManager) TodoItem() service.ITodoItemService {
	sm.todoItemRunner.Do(func() {
		sm.todoItem = service2.InitTodoItemService(sm.manager)
	})
	return sm.todoItem
}
