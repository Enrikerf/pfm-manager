package StreamResults

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type Query struct {
	BatchUuid Result.BatchId
	LastId    Result.Id
}
