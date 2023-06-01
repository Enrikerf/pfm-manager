package GetTaskBatches

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
)

type UseCase interface {
	List(query Query) ([]Result.Batch, error)
}

func New(
	findTaskBatches Repository.FindTaskBatches,
) UseCase {
	return &useCase{
		findTaskBatches: findTaskBatches,
	}
}

type useCase struct {
	findTaskBatches Repository.FindTaskBatches
}

func (useCase *useCase) List(query Query) ([]Result.Batch, error) {
	return useCase.findTaskBatches.Find(query.TaskId)
}
