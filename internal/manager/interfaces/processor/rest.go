package processor

import (
	"todo/internal/manager/interfaces/processor/rest"
)

type IRESTProcessor interface {
	Handler() rest.IHandlerManager
	Controllers() rest.IControllersManager
	Middlewares() rest.IMiddlewaresManager
}
