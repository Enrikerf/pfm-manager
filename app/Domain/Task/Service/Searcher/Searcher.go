package Searcher

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
)

type Searcher struct {
	FindByRepository Repository.FindBy
}

func (searcher *Searcher) Search(conditions interface{}) ([]Task.Task, error) {
	task, err := searcher.FindByRepository.FindBy(conditions)
	if err != nil {
		return nil, err
	}
	return task, nil
}
