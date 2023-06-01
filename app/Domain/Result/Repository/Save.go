package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
)

type Save interface {
	Persist(result Result.Result)
}
