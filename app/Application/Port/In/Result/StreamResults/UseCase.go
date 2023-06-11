package StreamResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type UseCase interface {
	Stream(query Query) (results []Result.Result, err error)
}

func New(
	findBatchPort Repository.FindBatch,
	findTaskRepository TaskRepository.Find,
	findBatchResultsAfterResult Repository.FindBatchResultsAfterResult,
) UseCase {
	return &useCase{
		FindBatchPort:               findBatchPort,
		TaskFinder:                  Finder.Finder{FindRepository: findTaskRepository},
		FindBatchResultsAfterResult: findBatchResultsAfterResult,
	}
}

type useCase struct {
	FindBatchPort               Repository.FindBatch
	TaskFinder                  Finder.Finder
	FindBatchResultsAfterResult Repository.FindBatchResultsAfterResult
}

func (service *useCase) Stream(query Query) (results []Result.Result, err error) {
	results, err = service.FindBatchResultsAfterResult.Find(query.BatchUuid, query.LastId)
	if err != nil {
		return nil, err
	}
	if results == nil {
		batch, err := service.FindBatchPort.Find(query.BatchUuid)
		if err != nil {
			return nil, err
		}
		task, err := service.TaskFinder.Find(batch.GetTaskId())
		if err != nil {
			return nil, err
		}
		if task.GetStatus().Value() != Status.Running {
			return nil, NewEndOfStreamError()
		}
	}
	return results, nil
}
