package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type DeleteMock struct {
	err error
}

func (mock DeleteMock) Delete(id Task.Id) error {
	return mock.err
}
