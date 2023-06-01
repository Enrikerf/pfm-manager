package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Creator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type UseCase interface {
	Create(command Command) (Task.Task, error)
}

func New(saveRepository Repository.Save, dispatcher Event.Dispatcher) UseCase {
	return &useCase{Creator.Creator{
		SaveRepository: saveRepository,
		Dispatcher:     dispatcher,
	}}
}

type useCase struct {
	creator Creator.Creator
}

func (useCase *useCase) Create(command Command) (Task.Task, error) {
	host, err := Host.NewVo(command.Host)
	if err != nil {
		return nil, err
	}
	port, err := Port.NewVo(command.Port)
	if err != nil {
		return nil, err
	}
	communicationMode, err := CommunicationMode.FromString(command.CommunicationMode)
	if err != nil {
		return nil, err
	}
	executionMode, err := ExecutionMode.FromString(command.ExecutionMode)
	if err != nil {
		return nil, err
	}
	var stepVos []Step.Vo
	for _, commandSentence := range command.CommandSentences {
		stepVo, err := Step.NewVo(commandSentence)
		if err != nil {
			return nil, err
		}
		stepVos = append(stepVos, stepVo)
	}
	if err != nil {
		return nil, err
	}
	return useCase.creator.Create(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
}
