package GetBatchResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
)

type UseCase interface {
	List(query Query) []Result.Result
}

func New(
	FindBatchResults Repository.FindBatchResults,
) UseCase {
	return &useCase{
		FindBatchResults,
	}
}

type useCase struct {
	FindBatchResults Repository.FindBatchResults
}

func (service *useCase) List(query Query) []Result.Result {
	results, _ := service.FindBatchResults.Find(query.BatchId)
	return results
}
