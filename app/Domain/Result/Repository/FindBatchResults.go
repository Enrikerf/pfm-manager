package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type FindBatchResults interface {
	Find(id Result.BatchId) ([]Result.Result,error)
}
