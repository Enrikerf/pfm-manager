package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type SaveBatch interface {
	Persist(result Result.Batch)
}
