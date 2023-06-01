package ReadTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/google/uuid"
)

type UseCase interface {
	Read(query Query) (Task.Task, error)
}

func New(findRepository Repository.Find) UseCase {
	return &useCase{Finder.Finder{FindRepository: findRepository}}
}

type useCase struct {
	finder Finder.Finder
}

func (useCase useCase) Read(query Query) (Task.Task, error) {
	uuidToFind, err := uuid.Parse(query.Uuid)
	if err != nil {
		return nil, errors.New("can't parse uuid")
	}
	task, err := useCase.finder.Find(Task.LoadId(uuidToFind))
	if err != nil {
		return nil, err
	}
	return task, nil
}
