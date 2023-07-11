package rest

import "todo/internal/manager/interfaces/processor/rest/controllers"

type IControllersManager interface {
	TodoItem() controllers.ITodoItemController
}
