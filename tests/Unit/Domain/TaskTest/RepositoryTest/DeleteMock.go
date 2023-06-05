package RepositoryTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

//go:generate mockery
type Delete interface {
	Delete(id Task.Id) error
}

type deleteMock struct {
	err error
}

func (receiver deleteMock) Delete(id Task.Id) error {
	return receiver.err

}

func BuildDeleteMock(err error) Repository.Delete {
	return &deleteMock{err}
}
