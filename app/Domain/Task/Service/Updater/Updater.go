package Updater

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	TaskEvent "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type Updater struct {
	FindRepository Repository.Find
	SaveRepository Repository.Save
	Dispatcher     Event.Dispatcher
}

func (updater *Updater) Update(
	id Task.Id,
	host Host.Vo,
	port Port.Vo,
	status Status.Status,
) error {

	task, err := updater.FindRepository.Find(id)
	if err != nil {
		return NewTaskNotFoundError()
	}
	if host != nil {
		task.SetHost(host)
	}
	if port != nil {
		task.SetPort(port)
	}
	if status != nil {
		task.SetStatus(status)
	}
	updater.SaveRepository.Persist(task)

	go updater.Dispatcher.Dispatch(TaskEvent.NewTaskUpdated(task.GetId().GetUuidString()))

	return nil
}
