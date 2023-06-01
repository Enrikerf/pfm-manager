package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type FindBatchResultsAfterResult interface {
	Find(id Result.BatchId, resultId Result.Id) ([]Result.Result, error)
}
