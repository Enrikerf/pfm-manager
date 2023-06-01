package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type FindTaskBatches interface {
	Find(id Task.Id) ([]Result.Batch, error)
}
