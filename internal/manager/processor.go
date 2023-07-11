package manager

import (
	"todo/internal/manager/interfaces"
	IProcessor "todo/internal/manager/interfaces/processor"

	"sync"
	"todo/internal/processor/rest"
)

type ProcessorManager struct {
	manager interfaces.IManager

	restProcessor       IProcessor.IRESTProcessor
	restProcessorRunner sync.Once
}

func InitProcessor(manager interfaces.IManager) *ProcessorManager {
	return &ProcessorManager{
		manager: manager,
	}

}

func (p *ProcessorManager) REST() IProcessor.IRESTProcessor {
	p.restProcessorRunner.Do(func() {
		p.restProcessor = rest.Init(p.manager)
	})
	return p.restProcessor
}
