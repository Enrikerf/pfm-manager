package CommunicateTaskManually

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/ManualTaskExecutor"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type UseCase interface {
	Communicate(command Command) (Result.Batch, error)
}

func New(
	manualTaskExecutor ManualTaskExecutor.Execute,
) UseCase {
	return &useCase{
		manualTaskExecutor: manualTaskExecutor,
	}
}

type useCase struct {
	manualTaskExecutor ManualTaskExecutor.Execute
}

func (service *useCase) Communicate(command Command) (Result.Batch, error) {
	taskId, err := Task.LoadIdFromString(command.TaskUuid)
	if err != nil {
		return nil, err
	}
	batch, err := service.manualTaskExecutor.Execute(taskId)
	if err != nil {
		return nil, err
	}
	return batch, nil
}
