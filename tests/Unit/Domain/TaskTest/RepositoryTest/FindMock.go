package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type findMock struct {
	task Task.Task
	err  error
}

func (receiver findMock) Find(id Task.Id) (Task.Task, error) {
	return receiver.task, receiver.err

}

func BuildFindMock(task Task.Task, err error) Repository.Find {
	return &findMock{task, err}
}
