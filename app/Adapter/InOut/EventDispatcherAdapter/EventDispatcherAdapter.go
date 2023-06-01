package EventDispatcherAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/TaskEventHandler"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	TaskEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
)

type EventDispatcherAdapter interface {
	Dispatch(event Event.Event)
}

type eventDispatcherAdapter struct {
	taskEventHandler TaskEventHandler.UseCase
}

func New(
	taskEventHandler TaskEventHandler.UseCase,
) EventDispatcherAdapter {
	self := &eventDispatcherAdapter{
		taskEventHandler,
	}
	return self
}

func (e *eventDispatcherAdapter) Dispatch(event Event.Event) {
	switch event.GetName() {
	case TaskEvent.TaskCreatedEventName:
		go e.taskEventHandler.Handle(event)
	}
}
