package runner

import (
	"todo/internal/manager/interfaces"
	controllers2 "todo/internal/manager/interfaces/processor/rest/controllers"
	"todo/internal/processor/rest/controllers"
)

type ControllersManager struct {
	todoItemController controllers2.ITodoItemController
}

func InitControllers(manager interfaces.IManager) *ControllersManager {
	return &ControllersManager{
		todoItemController: controllers.InitTodoItemController(manager),
	}
}

func (cm *ControllersManager) TodoItem() controllers2.ITodoItemController {
	return cm.todoItemController
}
