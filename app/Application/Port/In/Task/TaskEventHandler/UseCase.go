package TaskEventHandler

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/Looper"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type UseCase interface {
	Handle(created Event.TaskCreated)
}

func New(finder Finder.Finder, looper Looper.Looper) UseCase {
	return &taskEventHandler{
		finder: finder,
		looper: looper,
	}
}

type taskEventHandler struct {
	finder Finder.Finder
	looper Looper.Looper
}

func (useCase *taskEventHandler) Handle(taskChanged Event.TaskCreated) {
	taskId, err := Task.LoadIdFromString(taskChanged.GetEntityId())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task, err := useCase.finder.Find(taskId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if task.GetExecutionMode().Value() != ExecutionMode.Automatic || task.GetStatus().Value() != Status.Pending {
		return
	}
	if !useCase.looper.IsEnabled() {
		useCase.looper.Enable()
	}
}
