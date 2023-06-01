package Finder

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type Finder struct {
	FindRepository Repository.Find
}

func (finder *Finder) Find(id Task.Id) (Task.Task, error) {
	task, err := finder.FindRepository.Find(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, Error.NewTaskNotFoundError()
	}
	return task, nil
}
