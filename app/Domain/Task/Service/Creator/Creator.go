package Creator

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	TaskEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Creator struct {
	SaveRepository Repository.Save
	Dispatcher     Event.Dispatcher
}

func (creator *Creator) Create(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) (Task.Task, error) {

	var task, err = Task.New(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
	if err != nil {
		return nil, err
	}
	creator.SaveRepository.Persist(task)
	creator.Dispatcher.Dispatch(TaskEvent.NewTaskCreated(task.GetId().GetUuidString()))

	return task, nil
}
