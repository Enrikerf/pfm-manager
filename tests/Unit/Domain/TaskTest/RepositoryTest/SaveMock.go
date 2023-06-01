package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type SaveMock struct {
}

func (mock SaveMock) Persist(task Task.Task) {

}
