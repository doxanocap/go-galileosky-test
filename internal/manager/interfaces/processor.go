package interfaces

import "todo/internal/manager/interfaces/processor"

type IProcessor interface {
	REST() processor.IRESTProcessor
}
