package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type saveMock struct {
}

func (s saveMock) Persist(task Task.Task) {

}

func BuildSaveMock() Repository.Save {
	return saveMock{}
}
