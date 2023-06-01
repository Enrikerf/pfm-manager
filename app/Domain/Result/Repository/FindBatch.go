package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type FindBatch interface {
	Find(id Result.BatchId) (Result.Batch, error)
}
