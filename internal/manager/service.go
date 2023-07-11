package manager

import (
	"todo/internal/manager/interfaces"
)

type ServiceManager struct {
	manager interfaces.IManager
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}
