package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type findByMock struct {
	tasks []Task.Task
	err   error
}

func (receiver findByMock) FindBy(conditions interface{}) ([]Task.Task, error) {
	return receiver.tasks, receiver.err

}

func BuildFindByMock(tasks []Task.Task, err error) Repository.FindBy {
	return &findByMock{tasks, err}
}
