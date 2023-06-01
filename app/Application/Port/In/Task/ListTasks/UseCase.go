package ListTasks

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Searcher"
)

type UseCase interface {
	List(query Query) ([]Task.Task, error)
}

func New(findByRepository Repository.FindBy) UseCase {
	return &useCase{Searcher.Searcher{FindByRepository: findByRepository}}
}

type useCase struct {
	searcher Searcher.Searcher
}

func (useCase useCase) List(query Query) ([]Task.Task, error) {
	tasks, err := useCase.searcher.Search(query)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
