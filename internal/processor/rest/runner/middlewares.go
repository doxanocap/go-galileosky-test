package runner

import (
	"todo/internal/manager/interfaces"
	"todo/internal/manager/interfaces/processor/rest"
)

type Middlewares struct {
}

func InitMiddlewares(manager interfaces.IManager) rest.IMiddlewaresManager {
	return &Middlewares{}
}
