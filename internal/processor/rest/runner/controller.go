package runner

import (
	"todo/internal/manager/interfaces"
)

type ControllersManager struct {
}

func InitControllers(manager interfaces.IManager) *ControllersManager {
	return &ControllersManager{}
}
